package parsers

import (
	"encoding/json"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"io/ioutil"
	"strings"
)

type qodanaSource struct {
	Path   string
	Line   int64
	Offset int64
}

type qodanaProblem struct {
	Severity    string
	Comment     string
	DetailsInfo string
	Sources     []qodanaSource
}

type qodanaReport struct {
	Version     string
	ListProblem []qodanaProblem
}

func qodanaSeverityToAnnotationLevel(severity string) string {
	switch strings.ToLower(severity) {
	case "critical":
		return "failure"
	case "high":
		fallthrough
	case "moderate":
		return "warning"
	default:
		return "notice"
	}
}

func ParseQodanaAnnotations(path string) (error, []model.Annotation) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}

	var report qodanaReport
	err = json.Unmarshal(data, &report)
	if err != nil {
		return err, nil
	}

	result := make([]model.Annotation, 0)

	for _, problem := range report.ListProblem {
		for _, source := range problem.Sources {
			var parsedAnnotation = model.Annotation{
				Level:      qodanaSeverityToAnnotationLevel(problem.Severity),
				Message:    problem.Comment,
				RawDetails: problem.DetailsInfo,
				Location: &model.FileLocation{
					Path:        source.Path,
					StartLine:   source.Line,
					EndLine:     source.Line,
					StartColumn: source.Offset,
					EndColumn:   source.Offset,
				},
			}

			result = append(result, parsedAnnotation)
		}
	}

	return nil, result
}
