package parsers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/cirruslabs/cirrus-ci-annotations/model"
)

const supportedFormatVersion = 6

// severityMapping maps Android Lint severity to Github's annotation level
// https://android.googlesource.com/platform/tools/base/+/studio-master-dev/lint/libs/lint-api/src/main/java/com/android/tools/lint/detector/api/Severity.kt
// we leave out the Ignore severity type as we don't want to report it
var severityMapping = map[string]model.AnnotationLevel{
	"Fatal":       model.LevelFailure,
	"Error":       model.LevelFailure,
	"Warning":     model.LevelWarning,
	"Information": model.LevelNotice,
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

	if report.Format > supportedFormatVersion {
		return fmt.Errorf("Unsupported report version %d. Maximum supported version is %d", report.Format, supportedFormatVersion), nil
	}

	result := make([]model.Annotation, 0)
	for _, issue := range report.Issues {

		// we skip the Ignore severity as there is no mapping for it
		level, found := severityMapping[issue.Severity]

		if found {
			for _, location := range issue.Locations {
				var parsedAnnotation = model.Annotation{
					Level:       level,
					Message:     issue.Message,
					RawDetails:  issue.Explanation,
					Path:        location.Filename,
					StartLine:   location.Line,
					EndLine:     location.Line,
					StartColumn: location.Column,
					EndColumn:   location.Column,
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
