package parsers_test

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/cirruslabs/cirrus-ci-annotations/parsers"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestParseCirrusAnnotations(t *testing.T) {
	err, annotations := parsers.ParseCirrusAnnotations(filepath.Join("..", "testdata", "json", "cirrus.json"))
	if err != nil {
		t.Fatal(err)
	}

	expected := []model.Annotation{
		{
			Level:       model.LevelFailure,
			Message:     "some message",
			RawDetails:  "some lengthy details",
			Path:        "main.go",
			StartLine:   10,
			EndLine:     10,
			StartColumn: 5,
			EndColumn:   5,
		},
	}

	assert.Equal(t, expected, annotations)
}
