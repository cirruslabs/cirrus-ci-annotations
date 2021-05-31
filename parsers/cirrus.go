package parsers

import (
	"encoding/json"
	"fmt"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"os"
)

const currentReportVersion = "1"

type cirrusReport struct {
	Version     string             `json:"version"`
	Annotations []model.Annotation `json:"annotations"`
}

func ParseCirrusAnnotations(path string) (error, []model.Annotation) {
	reportFile, err := os.Open(path)
	if err != nil {
		return err, nil
	}
	defer reportFile.Close()

	decoder := json.NewDecoder(reportFile)

	var report cirrusReport

	if err := decoder.Decode(&report); err != nil {
		return err, nil
	}

	// Hard version checks for easy backwards compatibility in the future
	if report.Version == "" {
		return fmt.Errorf("you should specify a report's version, currently supported version is %q",
			currentReportVersion), nil
	}
	if report.Version != currentReportVersion {
		return fmt.Errorf("unsupported report version %q, currently only version %q is supported",
			report.Version, currentReportVersion), nil
	}

	return nil, report.Annotations
}
