package parsers

import (
	"encoding/json"
	"fmt"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"io/ioutil"
)

type GoLangCIReport struct {
	Issues []GoLangCIIssue
}

type GoLangCIIssue struct {
	FromLinter  string
	Text        string
	SourceLines []string
	Pos         GoLangCIPosition
}

type GoLangCIPosition struct {
	Filename string
	Offset   int64
	Line     int64
	Column   int64
}

func ParseGoLangCIAnnotations(path string) (error, []model.Annotation) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}
	var report GoLangCIReport
	err = json.Unmarshal(data, &report)
	if err != nil {
		return err, nil
	}
	result := make([]model.Annotation, 0)
	for _, issue := range report.Issues {
		var parsedAnnotation = model.Annotation{
			Type:               model.TestResultAnnotationType,
			Level:              "failure",
			FullyQualifiedName: issue.FromLinter,
			Message:            fmt.Sprintf("%s (%s)", issue.Text, issue.FromLinter),
			Location: &model.FileLocation{
				Path:        issue.Pos.Filename,
				StartLine:   issue.Pos.Line,
				EndLine:     issue.Pos.Line,
				StartColumn: issue.Pos.Column,
			},
		}

		result = append(
			result,
			parsedAnnotation,
		)
	}
	return nil, result
}
