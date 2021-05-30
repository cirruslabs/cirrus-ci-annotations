package parsers_test

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/cirruslabs/cirrus-ci-annotations/parsers"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestXclogparser(t *testing.T) {
	expected := []model.Annotation{
		{
			Level:   model.LevelFailure,
			Message: "Cannot find 'printeh' in scope",
			RawDetails: `/var/folders/sr/b58hwhtj0jbcf4r09zmg0wlc0000gn/T/cirrus-ci-build/noapp/main.swift:3:1: error: cannot find 'printeh' in scope
printeh("Hello, World!")
^~~~~~~`,
			Location: &model.FileLocation{
				Path:        "/var/folders/sr/b58hwhtj0jbcf4r09zmg0wlc0000gn/T/cirrus-ci-build/noapp/main.swift",
				StartLine:   3,
				EndLine:     3,
				StartColumn: 1,
				EndColumn:   1,
			},
		},
	}

	err, actual := parsers.ParseXclogparserAnnotations(filepath.Join("..", "testdata", "json", "xclogparser.json"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
