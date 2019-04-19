package model

type AnnotationType int32

const (
	GenericAnnotationType    AnnotationType = 0
	TestResultAnnotationType AnnotationType = 1
	LintResultAnnotationType AnnotationType = 2
	AnalysysAnnotationType   AnnotationType = 3
)

type FileLocation struct {
	Path        string
	StartLine   int64
	EndLine     int64
	StartColumn int64
	EndColumn   int64
}

// mimics https://developer.github.com/v3/checks/runs/#annotations-object
type Annotation struct {
	Type               AnnotationType
	Level              string
	Message            string
	RawDetails         string
	FullyQualifiedName string
	Location           *FileLocation
}
