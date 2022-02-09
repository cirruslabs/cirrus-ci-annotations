package parsers

import (
	"encoding/json"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"os"
)

type bufEntry struct {
	Path        string `json:"path"`
	StartLine   int64  `json:"start_line"`
	EndLine     int64  `json:"end_line"`
	StartColumn int64  `json:"start_column"`
	EndColumn   int64  `json:"end_column"`
	Type        string `json:"type"`
	Message     string `json:"message"`
}

func ParseBufAnnotations(path string) (error, []model.Annotation) {
	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}

	var result []model.Annotation

	decoder := json.NewDecoder(file)

	for decoder.More() {
		var entry bufEntry

		if err := decoder.Decode(&entry); err != nil {
			return err, nil
		}

		result = append(result, model.Annotation{
			Level:       model.LevelWarning,
			Message:     entry.Message,
			Path:        entry.Path,
			StartLine:   entry.StartLine,
			EndLine:     entry.EndLine,
			StartColumn: entry.StartColumn,
			EndColumn:   entry.EndColumn,
		})
	}
	return nil, result
}
