package model

type FileLocation struct {
	Path        string
	StartLine   int64
	EndLine     int64
	StartColumn int64
	EndColumn   int64
}

// mimics https://developer.github.com/v3/checks/runs/#annotations-object
type Annotation struct {
	Level              string
	Message            string
	RawDetails         string
	Location           *FileLocation
}
