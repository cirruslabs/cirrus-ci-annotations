package parsers

import (
	"encoding/json"
	"fmt"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"os"
)

type swiftFormatEntry struct {
	File   string `json:"file"`
	Line   int64  `json:"line"`
	Reason string `json:"reason"`
	RuleID string `json:"rule_id"`
}

func ParseSwiftFormatAnnotations(path string) (error, []model.Annotation) {
	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}

	var result []model.Annotation

	decoder := json.NewDecoder(file)

	for decoder.More() {
		var entries []swiftFormatEntry

		if err := decoder.Decode(&entries); err != nil {
			return err, nil
		}

		for _, entry := range entries {
			result = append(result, model.Annotation{
				Level:     model.LevelWarning,
				Message:   fmt.Sprintf("%s: %s", entry.RuleID, entry.Reason),
				Path:      entry.File,
				StartLine: entry.Line,
				EndLine:   entry.Line,
			})
		}
	}

	return nil, result
}
