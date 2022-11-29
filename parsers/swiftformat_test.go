package parsers

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestParseSwiftFormatAnnotations(t *testing.T) {
	expected := []model.Annotation{
		{
			Level:     model.LevelWarning,
			Message:   "indent: Indent code in accordance with the scope level.",
			Path:      "/tmp/working-directory/tart/Sources/tart/Prunable.swift",
			StartLine: 4,
			EndLine:   4,
		},
		{
			Level:     model.LevelWarning,
			Message:   "wrapArguments: Align wrapped function arguments or collection elements.",
			Path:      "/tmp/working-directory/tart/Passphrase/Words.swift",
			StartLine: 2052,
			EndLine:   2052,
		},
	}

	fixturePath := filepath.Join("..", "testdata", "json", "swiftformat.json")

	err, actual := ParseSwiftFormatAnnotations(fixturePath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
