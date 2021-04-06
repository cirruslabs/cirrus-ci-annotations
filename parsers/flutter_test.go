package parsers_test

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/cirruslabs/cirrus-ci-annotations/parsers"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestFlutterSucceeding(t *testing.T) {
	err, actual := parsers.ParseFlutterAnnotations(filepath.Join("..", "testdata", "json", "flutter-succeeding.json"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Empty(t, actual)
}

func TestFlutterFailing(t *testing.T) {
	expected := []model.Annotation{
		{
			Type:    model.TestResultAnnotationType,
			Level:   "warning",
			Message: "Counter value should start at 0",
			Location: &model.FileLocation{
				Path:        "/tmp/cirrus-ci-build/test/counter_test.dart",
				StartLine:   6,
				EndLine:     6,
				StartColumn: 5,
				EndColumn:   5,
			},
		},
		{
			Type:    model.TestResultAnnotationType,
			Level:   "warning",
			Message: "Counter value should be incremented",
			Location: &model.FileLocation{
				Path:        "/tmp/cirrus-ci-build/test/counter_test.dart",
				StartLine:   10,
				EndLine:     10,
				StartColumn: 5,
				EndColumn:   5,
			},
		},
		{
			Type:    model.TestResultAnnotationType,
			Level:   "warning",
			Message: "Counter value should be decremented",
			Location: &model.FileLocation{
				Path:        "/tmp/cirrus-ci-build/test/counter_test.dart",
				StartLine:   18,
				EndLine:     18,
				StartColumn: 5,
				EndColumn:   5,
			},
		},
	}

	err, actual := parsers.ParseFlutterAnnotations(filepath.Join("..", "testdata", "json", "flutter-failing.json"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
