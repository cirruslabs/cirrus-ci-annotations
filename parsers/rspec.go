package parsers

import (
	"encoding/json"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"io/ioutil"
)

type rspecException struct {
	Message string
}

type rspecExample struct {
	ID              string
	Status          string
	FullDescription string `json:"full_description"`
	FilePath        string `json:"file_path"`
	LineNumber      int64  `json:"line_number"`

	// Used when status is "pending"
	PendingMessage string `json:"pending_message"`

	// Used when status is "failed"
	Exception rspecException
}

type rspecReport struct {
	Examples []rspecExample
}

// ParseRSpecAnnotations parses RSpec's JSON output into annotations according to the "spec" ranged from 3.0.0[1] to 3.9.2[2].
// [1]: https://github.com/rspec/rspec-core/blob/v3.0.0/lib/rspec/core/example.rb
// [2]: https://github.com/rspec/rspec-core/blob/v3.9.2/lib/rspec/core/example.rb
func ParseRSpecAnnotations(path string) (error, []model.Annotation) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}

	var report rspecReport
	err = json.Unmarshal(data, &report)
	if err != nil {
		return err, nil
	}

	result := make([]model.Annotation, 0)

	for _, example := range report.Examples {
		var level, rawDetails string

		// Skip "passed" (which only adds noise) and deal with the rest of the states
		switch example.Status {
		case "pending":
			level = "notice"
			rawDetails = example.PendingMessage
		case "failed":
			level = "failure"
			rawDetails = example.Exception.Message
		default:
			continue
		}

		var parsedAnnotation = model.Annotation{
			Type:               model.TestResultAnnotationType,
			Level:              level,
			Message:            example.FullDescription,
			RawDetails:         rawDetails,
			FullyQualifiedName: example.ID,
			Location: &model.FileLocation{
				Path:      example.FilePath,
				StartLine: example.LineNumber,
				EndLine:   example.LineNumber,
			},
		}

		result = append(result, parsedAnnotation)
	}

	return nil, result
}
