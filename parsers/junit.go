package parsers

import (
	"fmt"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/cirruslabs/cirrus-ci-annotations/util"
	"github.com/cirruslabs/go-junit"
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
					Level:   model.LevelNotice,
					Message: fqn,
				}
			case junit.StatusFailed:
				path, startLine, endLine := util.GuessLocationIgnored(
					test.Error.Error(),
					[]string{
						"junit",
						"kotlin",
					},
				)

				parsedAnnotation = model.Annotation{
					Level:      model.LevelFailure,
					Message:    fqn,
					RawDetails: test.Error.Error(),
					Path:       path,
					StartLine:  startLine,
					EndLine:    endLine,
				}
			}

			if test.Properties["file"] != "" {
				line, _ := strconv.Atoi(test.Properties["line"])
				parsedAnnotation.Path = test.Properties["file"]
				parsedAnnotation.StartLine = int64(line)
				parsedAnnotation.EndLine = int64(line)
			}

			result = append(
				result,
				parsedAnnotation,
			)
		}
	}
	return nil, result
}
