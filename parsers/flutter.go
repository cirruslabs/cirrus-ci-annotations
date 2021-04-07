package parsers

import (
	"encoding/json"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"io"
	"os"
	"strings"
)

type flutterTest struct {
	ID     int64
	Name   string
	URL    string
	Line   int64
	Column int64
}

type flutterEntry struct {
	Type   string
	TestID int64
	Result string
	Test   flutterTest
}

func annotationFromTest(test flutterTest) model.Annotation {
	return model.Annotation{
		Type:    model.TestResultAnnotationType,
		Level:   "failure",
		Message: test.Name,
		Location: &model.FileLocation{
			Path:        strings.TrimPrefix(test.URL, "file://"),
			StartLine:   test.Line,
			EndLine:     test.Line,
			StartColumn: test.Column,
			EndColumn:   test.Column,
		},
	}
}

func ParseFlutterAnnotations(path string) (error, []model.Annotation) {
	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}

	decoder := json.NewDecoder(file)
	runningTests := map[int64]flutterTest{}
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
		switch entry.Type {
		case "testStart":
			runningTests[entry.Test.ID] = entry.Test
		case "testDone":
			if entry.Result != "success" {
				test, ok := runningTests[entry.TestID]
				if ok {
					result = append(result, annotationFromTest(test))
				}
			}
		}
	}

	return nil, result
}
