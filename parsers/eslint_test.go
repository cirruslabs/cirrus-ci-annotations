package parsers

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestParseESLintAnnotations(t *testing.T) {
	expected := []model.Annotation{
		{
			Level:       model.LevelFailure,
			Message:     "curly: Expected { after 'if' condition.",
			Path:        "path/to/file.js",
			StartLine:   2,
			EndLine:     2,
			StartColumn: 1,
			EndColumn:   1,
		},
		{
			Level:       model.LevelFailure,
			Message:     "no-process-exit: Don't use process.exit(); throw an error instead.",
			Path:        "path/to/file.js",
			StartLine:   3,
			EndLine:     3,
			StartColumn: 1,
			EndColumn:   1,
		},
	}

	// Fixture taken from "Working with Custom Formatters"[1] and converted to JSON
	// [1]: https://eslint.org/docs/developer-guide/working-with-custom-formatters#the-results-object
	err, actual := ParseESLintAnnotations(filepath.Join("..", "testdata", "eslint.json"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
