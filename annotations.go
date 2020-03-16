package annotations

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/cirruslabs/cirrus-ci-annotations/parsers"
	"os"
	"path/filepath"
)

func ParseAnnotations(format string, path string) (error, []model.Annotation) {
	switch format {
	case "junit":
		return parsers.ParseJUnitAnnotations(path)
	default:
		return nil, make([]model.Annotation, 0)
	}
}

// Makes sure that locations has validate relative to workDirPath path
func ValidateAnnotations(workDirPath string, annotations []model.Annotation) error {
	fileIndex := make(map[string]string)
	err := filepath.Walk(workDirPath, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			fileIndex[filepath.Base(path)], _ = filepath.Rel(workDirPath, path)
		}
		return nil
	})

	if err != nil {
		return err
	}

	for _, annotation := range annotations {
		location := annotation.Location
		if location == nil {
			continue
		}
		path := location.Path
		if filepath.IsAbs(path) {
			path, _ = filepath.Rel(workDirPath, path)
		}
		if _, err := os.Stat(filepath.Join(workDirPath, path)); os.IsNotExist(err) {
			path = fileIndex[filepath.Base(path)]
		}
		if path != "" {
			location.Path = path
		}
	}
	return nil
}
