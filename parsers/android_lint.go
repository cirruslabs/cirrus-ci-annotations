package parsers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/cirruslabs/cirrus-ci-annotations/model"
)

const supportedFormatVersion = 5

// severityMapping maps Android Lint severity to Github's annotation level
// https://android.googlesource.com/platform/tools/base/+/studio-master-dev/lint/libs/lint-api/src/main/java/com/android/tools/lint/detector/api/Severity.kt
// we leave out the Ignore severity type as we don't want to report it
var severityMapping = map[string]string{
	"Fatal":       "fatal",
	"Error":       "fatal",
	"Warning":     "warning",
	"Information": "notice",
}

type androidLintReport struct {
	XMLName xml.Name           `xml:"issues"`
	Issues  []androidLintIssue `xml:"issue"`
	Format  int64              `xml:"format,attr"`
}

type androidLintIssue struct {
	XMLName     xml.Name   `xml:"issue"`
	ID          string     `xml:"id,attr"`
	Severity    string     `xml:"severity,attr"`
	Message     string     `xml:"message,attr"`
	Category    string     `xml:"category,attr"`
	Priority    string     `xml:"priority,attr"`
	Summary     string     `xml:"summary,attr"`
	Explanation string     `xml:"explanation,attr"`
	Locations   []location `xml:"location"`
}

type location struct {
	XMLName  xml.Name `xml:"location"`
	Filename string   `xml:"file,attr"`
	Line     int64    `xml:"line,attr"`
	Column   int64    `xml:"column,attr"`
}

func ParseAndroidLintAnnotations(path string) (error, []model.Annotation) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}
	var report androidLintReport
	err = xml.Unmarshal(data, &report)
	if err != nil {
		return err, nil
	}

	if report.Format != supportedFormatVersion {
		return fmt.Errorf("Unsupported report version %d. Supported version is %d", report.Format, supportedFormatVersion), nil
	}

	result := make([]model.Annotation, 0)
	for _, issue := range report.Issues {

		// we skip the Ignore severity as there is no mapping for it
		level, found := severityMapping[issue.Severity]

		if found {
			for _, location := range issue.Locations {
				var parsedAnnotation = model.Annotation{
					Type:               model.LintResultAnnotationType,
					Level:              level,
					FullyQualifiedName: issue.ID,
					Message:            issue.Message,
					RawDetails:         issue.Explanation,
					Location: &model.FileLocation{
						Path:        location.Filename,
						StartLine:   location.Line,
						EndLine:     location.Line,
						StartColumn: location.Column,
						EndColumn:   location.Column,
					},
				}

				result = append(
					result,
					parsedAnnotation,
				)
			}
		}
	}
	return nil, result
}
