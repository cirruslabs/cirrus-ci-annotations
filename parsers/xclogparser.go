package parsers

import (
	"encoding/json"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"io/ioutil"
	"strings"
)

type xclogparserEntry struct {
	Title                string
	Detail               string
	DocumentURL          string
	StartingLineNumber   int64
	EndingLineNumber     int64
	StartingColumnNumber int64
	EndingColumnNumber   int64
}

type xclogparserReport struct {
	Errors   []xclogparserEntry
	Warnings []xclogparserEntry
	Notes    []xclogparserEntry
}

func (entry *xclogparserEntry) ToAnnotation(level model.AnnotationLevel) model.Annotation {
	return model.Annotation{
		Level:      level,
		Message:    entry.Title,
		RawDetails: entry.Detail,
		Location: &model.FileLocation{
			Path:        strings.TrimPrefix(entry.DocumentURL, "file://"),
			StartLine:   entry.StartingLineNumber,
			EndLine:     entry.EndingLineNumber,
			StartColumn: entry.StartingColumnNumber,
			EndColumn:   entry.EndingColumnNumber,
		},
	}
}

func ParseXclogparserAnnotations(path string) (error, []model.Annotation) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}

	var reports []xclogparserReport
	err = json.Unmarshal(data, &reports)
	if err != nil {
		return err, nil
	}

	result := make([]model.Annotation, 0)

	for _, report := range reports {
		for _, reportError := range report.Errors {
			if reportError.DocumentURL == "" {
				continue
			}

			result = append(result, reportError.ToAnnotation(model.LevelFailure))
		}

		for _, reportWarning := range report.Warnings {
			if reportWarning.DocumentURL == "" {
				continue
			}

			result = append(result, reportWarning.ToAnnotation(model.LevelWarning))
		}

		for _, reportNote := range report.Notes {
			if reportNote.DocumentURL == "" {
				continue
			}

			result = append(result, reportNote.ToAnnotation(model.LevelNotice))
		}
	}

	return nil, result
}
