package parsers_test

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/cirruslabs/cirrus-ci-annotations/parsers"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestQodana(t *testing.T) {
	expected := []*model.Annotation{
		{
			Level:   model.LevelFailure,
			Message: "Cannot resolve symbol 'ComponentSelection' ",
			RawDetails: `<html>
<body>
Allows you to see problems reported by language annotators in the results of batch code inspection.
</body>
</html>`,
			Path:        "build.gradle",
			StartLine:   37,
			EndLine:     37,
			StartColumn: 16,
			EndColumn:   16,
		},
		{
			Level:   model.LevelFailure,
			Message: "Cannot resolve symbol 'FileType' ",
			RawDetails: `<html>
<body>
Allows you to see problems reported by language annotators in the results of batch code inspection.
</body>
</html>`,
			Path:        "settings.gradle",
			StartLine:   3,
			EndLine:     3,
			StartColumn: 24,
			EndColumn:   24,
		},
		{
			Level:   model.LevelFailure,
			Message: "Cannot resolve symbol 'FileVisitResult' ",
			RawDetails: `<html>
<body>
Allows you to see problems reported by language annotators in the results of batch code inspection.
</body>
</html>`,
			Path:        "settings.gradle",
			StartLine:   4,
			EndLine:     4,
			StartColumn: 24,
			EndColumn:   24,
		},
		{
			Level:   model.LevelFailure,
			Message: "Cannot resolve symbol 'FileVisitResult' ",
			RawDetails: `<html>
<body>
Allows you to see problems reported by language annotators in the results of batch code inspection.
</body>
</html>`,
			Path:        "settings.gradle",
			StartLine:   5,
			EndLine:     5,
			StartColumn: 24,
			EndColumn:   24,
		},
		{
			Level:   model.LevelWarning,
			Message: "Kotlin version that is used for building with Gradle (1.3.72) differs from the one bundled into the IDE plugin (1.4.10)",
			RawDetails: `<html>
<body>
This inspection reports that different IDE and Gradle plugin versions are used.
This can cause inconsistencies between IDE and Gradle builds in error reporting or code behaviour.
</body>
</html>`,
			Path:        "build.gradle",
			StartLine:   11,
			EndLine:     11,
			StartColumn: 4,
			EndColumn:   4,
		},
	}

	err, actual := parsers.ParseQodanaAnnotations(filepath.Join("..", "testdata", "json", "qodana.json"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
