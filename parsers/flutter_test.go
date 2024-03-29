package parsers_test

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/cirruslabs/cirrus-ci-annotations/parsers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			Level:       model.LevelFailure,
			Message:     "Counter value should start at 0",
			RawDetails:  "Expected: <0>\n  Actual: <1>\n",
			Path:        "/tmp/cirrus-ci-build/test/counter_test.dart",
			StartLine:   6,
			EndLine:     6,
			StartColumn: 5,
			EndColumn:   5,
		},
		{
			Level:       model.LevelFailure,
			Message:     "Counter value should be incremented",
			RawDetails:  "Expected: <1>\n  Actual: <2>\n",
			Path:        "/tmp/cirrus-ci-build/test/counter_test.dart",
			StartLine:   10,
			EndLine:     10,
			StartColumn: 5,
			EndColumn:   5,
		},
		{
			Level:       model.LevelFailure,
			Message:     "Counter value should be decremented",
			RawDetails:  "Expected: <-1>\n  Actual: <0>\n",
			Path:        "/tmp/cirrus-ci-build/test/counter_test.dart",
			StartLine:   18,
			EndLine:     18,
			StartColumn: 5,
			EndColumn:   5,
		},
	}

	err, actual := parsers.ParseFlutterAnnotations(filepath.Join("..", "testdata", "json", "flutter-failing.json"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

// TestFlutterArrays makes sure that we ignore the events that Flutter
// started providing in the --machine output recently.
//
// https://github.com/cirruslabs/cirrus-ci-docs/issues/1245
func TestFlutterArrays(t *testing.T) {
	err, _ := parsers.ParseFlutterAnnotations(filepath.Join("..", "testdata", "json", "flutter-arrays.json"))
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err)
}
