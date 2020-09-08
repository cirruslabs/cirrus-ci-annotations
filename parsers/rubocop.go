package parsers

import (
	"encoding/json"
	"fmt"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"io/ioutil"
)

// rubocopSeverityMapping translates RuboCop's severity constants[1] into GitHub
// annotation levels using RuboCop's SimpleTextFormatter colors[2] as a hint.
// [1]: https://www.rubydoc.info/gems/rubocop/RuboCop/Cop/Severity
// [2]: https://www.rubydoc.info/gems/rubocop/RuboCop/Formatter/SimpleTextFormatter#COLOR_FOR_SEVERITY-constant
var rubocopSeverityMapping = map[string]string{
	"refactor":   "notice",
	"convention": "notice",
	"warning":    "warning",
	"error":      "failure",
	"fatal":      "failure",
}

type rubocopLocation struct {
	Line   int64
	Column int64
}

type rubocopOffense struct {
	Severity  string
	Message   string
	CopName   string `json:"cop_name"`
	Corrected bool
	Location  rubocopLocation
}

type rubocopFile struct {
	Path     string
	Offenses []rubocopOffense
}

type rubocopReport struct {
	Files []rubocopFile
}

// ParseRuboCopAnnotations parses RuboCop's JSON output into annotations according to the documentation
// ranged from 0.53.0[1] to 0.87.1[2].
// [1]: https://github.com/rubocop-hq/rubocop/blob/v0.53.0/manual/formatters.md#json-formatter
// [2]: https://github.com/rubocop-hq/rubocop/blob/v0.87.1/docs/modules/ROOT/pages/formatters.adoc#json-formatter
func ParseRuboCopAnnotations(path string) (error, []model.Annotation) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}

	var report rubocopReport
	err = json.Unmarshal(data, &report)
	if err != nil {
		return err, nil
	}

	result := make([]model.Annotation, 0)

	for _, file := range report.Files {
		for offenseNumber, offense := range file.Offenses {
			// Map RuboCop's severity to GitHub's annotation level,
			// skipping unknown severities.
			level, found := rubocopSeverityMapping[offense.Severity]
			if !found {
				continue
			}

			// No matter the severity, corrected offenses get demoted to "notice" level
			if offense.Corrected {
				level = "notice"
			}

			// RuboCop does not provide a FQN, so we craft one ourselves
			fqn := fmt.Sprintf("%s-%d", file.Path, offenseNumber)

			var parsedAnnotation = model.Annotation{
				Type:               model.LintResultAnnotationType,
				Level:              level,
				Message:            fmt.Sprintf("%s: %s", offense.CopName, offense.Message),
				FullyQualifiedName: fqn,
				Location: &model.FileLocation{
					Path:        file.Path,
					StartLine:   offense.Location.Line,
					EndLine:     offense.Location.Line,
					StartColumn: offense.Location.Column,
					EndColumn:   offense.Location.Column,
				},
			}

			result = append(result, parsedAnnotation)
		}
	}

	return nil, result
}
