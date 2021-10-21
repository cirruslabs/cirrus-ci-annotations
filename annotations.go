package annotations

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/cirruslabs/cirrus-ci-annotations/parsers"
)

func ParseAnnotations(format string, path string) (error, []model.Annotation) {
	switch strings.ToLower(format) {
	case "junit":
		return parsers.ParseJUnitAnnotations(path)
	case "golangci":
		return parsers.ParseGoLangCIAnnotations(path)
	case "android-lint":
		return parsers.ParseAndroidLintAnnotations(path)
	case "rspec":
		return parsers.ParseRSpecAnnotations(path)
	case "rubocop":
		return parsers.ParseRuboCopAnnotations(path)
	case "qodana":
		return parsers.ParseQodanaAnnotations(path)
	case "xclogparser":
		return parsers.ParseXclogparserAnnotations(path)
	case "flutter":
		return parsers.ParseFlutterAnnotations(path)
	case "cirrus":
		return parsers.ParseCirrusAnnotations(path)
	default:
		return nil, make([]model.Annotation, 0)
	}
}

// Makes sure that locations has validate relative to workDirPath path
func ValidateAnnotations(workDirPath string, annotations []model.Annotation) ([]model.Annotation, error) {
	var result []model.Annotation

	fileIndex := make(map[string]string)
	err := filepath.Walk(workDirPath, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			fileIndex[filepath.Base(path)], _ = filepath.Rel(workDirPath, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, annotation := range annotations {
		path := annotation.Path
		if filepath.IsAbs(path) {
			path, _ = filepath.Rel(workDirPath, path)
		}
		if _, err := os.Stat(filepath.Join(workDirPath, path)); os.IsNotExist(err) {
			path = fileIndex[filepath.Base(path)]
		}
		if path != "" {
			annotation.Path = path
		}

		result = append(result, annotation)
	}

	return result, nil
}
