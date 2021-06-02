package parsers

import (
	"encoding/json"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"os"
)

func ParseCirrusAnnotations(path string) (error, []model.Annotation) {
	reportFile, err := os.Open(path)
	if err != nil {
		return err, nil
	}
	defer reportFile.Close()

	var annotations []model.Annotation

	decoder := json.NewDecoder(reportFile)

	for decoder.More() {
		var annotation model.Annotation

		if err := decoder.Decode(&annotation); err != nil {
			return err, nil
		}

		annotations = append(annotations, annotation)
	}

	return nil, annotations
}
