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
	Error  string
	TestID int64
	Result string
	Test   flutterTest
}

func annotationFromTest(test *flutterTest) *model.Annotation {
	return &model.Annotation{
		Message:     test.Name,
		Path:        strings.TrimPrefix(test.URL, "file://"),
		StartLine:   test.Line,
		EndLine:     test.Line,
		StartColumn: test.Column,
		EndColumn:   test.Column,
	}
}

func ParseFlutterAnnotations(path string) (error, []model.Annotation) {
	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}

	decoder := json.NewDecoder(file)
	runningTests := map[int64]*model.Annotation{}
	result := make([]model.Annotation, 0)

	for {
		// Get the next raw-encoded JSON value
		var raw json.RawMessage

		if err := decoder.Decode(&raw); err != nil {
			if err == io.EOF {
				break
			}

			return err, nil
		}

		// Filter out JSON values that don't look like JSON objects
		if len(raw) > 0 && raw[0] != '{' {
			continue
		}

		// Parse "flutter test --machine" entry
		var entry flutterEntry

		if err := json.Unmarshal(raw, &entry); err != nil {
			return err, nil
		}

		// Does it look like a TestStartEvent[1] with a valid file associated with it?
		//
		// [1]: https://github.com/dart-lang/test/blob/master/pkgs/test/doc/json_reporter.schema.json
		switch entry.Type {
		case "testStart":
			runningTests[entry.Test.ID] = annotationFromTest(&entry.Test)
		case "error":
			if annotation, ok := runningTests[entry.TestID]; ok {
				annotation.RawDetails = entry.Error
			}
		case "testDone":
			if entry.Result != "success" {
				if annotation, ok := runningTests[entry.TestID]; ok {
					annotation.Level = model.LevelFailure
					result = append(result, *annotation)
					delete(runningTests, entry.TestID)
				}
			}
		}
	}

	return nil, result
}
