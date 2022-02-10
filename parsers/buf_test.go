package parsers

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestParseBufAnnotations(t *testing.T) {
	expected := []model.Annotation{
		{
			Level:       model.LevelWarning,
			Message:     "Package name \"google.type\" should be suffixed with a correctly formed version, such as \"google.type.v1\".",
			Path:        "google/type/datetime.proto",
			StartLine:   17,
			EndLine:     17,
			StartColumn: 1,
			EndColumn:   21,
		},
		{
			Level:       model.LevelWarning,
			Message:     "Field name \"petID\" should be lower_snake_case, such as \"pet_id\".",
			Path:        "pet/v1/pet.proto",
			StartLine:   42,
			EndLine:     42,
			StartColumn: 10,
			EndColumn:   15,
		},
	}

	// Fixture taken from "Run lint"[1] example
	// [1]: https://docs.buf.build/lint/usage#run-lint
	err, actual := ParseBufAnnotations(filepath.Join("..", "testdata", "json", "buf.json"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
