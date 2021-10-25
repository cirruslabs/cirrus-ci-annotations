package parsers

import (
	"encoding/json"
	"fmt"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"io/ioutil"
)

type eslintMessage struct {
	RuleID   string `json:"ruleId"`
	Severity int    `json:"severity"`
	Message  string `json:"message"`
	Line     int64  `json:"line"`
	Column   int64  `json:"column"`
	NodeType string `json:"nodeType"`
}

type eslintResult struct {
	FilePath string          `json:"filePath"`
	Messages []eslintMessage `json:"messages"`
}

func eslintSeverityToAnnotationLevel(eslintSeverity int) model.AnnotationLevel {
	switch eslintSeverity {
	case 1:
		return model.LevelWarning
	case 2:
		return model.LevelFailure
	default:
		return model.LevelNotice
	}
}

func ParseESLintAnnotations(path string) (error, []model.Annotation) {
	var result []model.Annotation

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}

	var eslintResults []eslintResult
	err = json.Unmarshal(data, &eslintResults)
	if err != nil {
		return err, nil
	}

	for _, eslintResult := range eslintResults {
		for _, eslintMessage := range eslintResult.Messages {
			result = append(result, model.Annotation{
				Level:       eslintSeverityToAnnotationLevel(eslintMessage.Severity),
				Message:     fmt.Sprintf("%s: %s", eslintMessage.RuleID, eslintMessage.Message),
				Path:        eslintResult.FilePath,
				StartLine:   eslintMessage.Line,
				EndLine:     eslintMessage.Line,
				StartColumn: eslintMessage.Column,
				EndColumn:   eslintMessage.Column,
			})
		}
	}

	return nil, result
}
