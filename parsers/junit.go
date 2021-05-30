package parsers

import (
	"fmt"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/cirruslabs/cirrus-ci-annotations/util"
	"github.com/joshdk/go-junit"
	"strconv"
)

func ParseJUnitAnnotations(path string) (error, []model.Annotation) {
	suites, err := junit.IngestFile(path)
	if err != nil {
		return err, nil
	}
	result := make([]model.Annotation, 0)
	for _, suite := range suites {
		for _, test := range suite.Tests {
			fqn := fmt.Sprintf("%s.%s", test.Classname, test.Name)
			var parsedAnnotation model.Annotation
			switch test.Status {
			case junit.StatusPassed:
				parsedAnnotation = model.Annotation{
					Level:              "notice",
					Message:            fqn,
				}
			case junit.StatusFailed:
				parsedAnnotation = model.Annotation{
					Level:              "failure",
					Message:            fqn,
					RawDetails:         test.Error.Error(),
					Location: util.GuessLocationIgnored(
						test.Error.Error(),
						[]string{
							"junit",
							"kotlin",
						},
					),
				}
			}

			if test.Properties["file"] != "" {
				line, _ := strconv.Atoi(test.Properties["line"])
				parsedAnnotation.Location = &model.FileLocation{
					Path:      test.Properties["file"],
					StartLine: int64(line),
					EndLine:   int64(line),
				}
			}

			result = append(
				result,
				parsedAnnotation,
			)
		}
	}
	return nil, result
}
