package model

import "fmt"

type AnnotationLevel int

const (
	LevelNotice AnnotationLevel = iota
	LevelWarning
	LevelFailure
)

func (al *AnnotationLevel) String() string {
	switch *al {
	case LevelNotice:
		return "notice"
	case LevelWarning:
		return "warning"
	case LevelFailure:
		return "failure"
	default:
		panic(fmt.Sprintf("unhandled annotation level: %d", *al))
	}
}

type FileLocation struct {
	Path        string
	StartLine   int64
	EndLine     int64
	StartColumn int64
	EndColumn   int64
}

// mimics https://developer.github.com/v3/checks/runs/#annotations-object
type Annotation struct {
	Level      AnnotationLevel
	Message    string
	RawDetails string
	Location   *FileLocation
}
