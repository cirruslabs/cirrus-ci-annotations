package model

import (
	"encoding/json"
	"fmt"
)

type AnnotationLevel int

const (
	literalNotice  = "notice"
	literalWarning = "warning"
	literalFailure = "failure"
)

const (
	LevelNotice AnnotationLevel = iota
	LevelWarning
	LevelFailure
)

func (al *AnnotationLevel) String() string {
	switch *al {
	case LevelNotice:
		return literalNotice
	case LevelWarning:
		return literalWarning
	case LevelFailure:
		return literalFailure
	default:
		panic(fmt.Sprintf("unhandled annotation level: %d", *al))
	}
}

func (al *AnnotationLevel) UnmarshalJSON(data []byte) error {
	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	switch value {
	case literalNotice:
		*al = LevelNotice
	case literalWarning:
		*al = LevelWarning
	case literalFailure:
		*al = LevelFailure
	default:
		return fmt.Errorf("unsupported annotation level: %q", value)
	}

	return nil
}

// mimics https://developer.github.com/v3/checks/runs/#annotations-object
type Annotation struct {
	Level      AnnotationLevel `json:"level"`
	Message    string          `json:"message"`
	RawDetails string          `json:"raw_details"`

	Path        string `json:"path"`
	StartLine   int64  `json:"start_line"`
	EndLine     int64  `json:"end_line"`
	StartColumn int64  `json:"start_column"`
	EndColumn   int64  `json:"end_column"`
}
