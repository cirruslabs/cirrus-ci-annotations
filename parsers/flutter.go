package parsers

import (
	"encoding/json"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"io"
	"os"
	"strings"
)

type flutterTest struct {
	Name   string
	URL    string
	Line   int64
	Column int64
}

type flutterEntry struct {
	Type string
	Test flutterTest
}

func ParseFlutterAnnotations(path string) (error, []model.Annotation) {
	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}

	decoder := json.NewDecoder(file)
	result := make([]model.Annotation, 0)

	for {
		var entry flutterEntry

		if err := decoder.Decode(&entry); err != nil {
			if err == io.EOF {
				break
			}

			return err, nil
		}

		// Does it look like a TestStartEvent[1] with a valid file associated with it?
		//
		// [1]: https://github.com/dart-lang/test/blob/master/pkgs/test/doc/json_reporter.schema.json
		if entry.Type != "testStart" || entry.Test.URL == "" {
			continue
		}

		var parsedAnnotation = model.Annotation{
			Type:    model.TestResultAnnotationType,
			Level:   "warning",
			Message: entry.Test.Name,
			Location: &model.FileLocation{
				Path:        strings.TrimPrefix(entry.Test.URL, "file://"),
				StartLine:   entry.Test.Line,
				EndLine:     entry.Test.Line,
				StartColumn: entry.Test.Column,
				EndColumn:   entry.Test.Column,
			},
		}

		result = append(result, parsedAnnotation)
	}

	return nil, result
}
