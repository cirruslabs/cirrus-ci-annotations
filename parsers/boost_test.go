package parsers

import (
	"path/filepath"
	"testing"

	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/go-test/deep"
)

func TestParseBoostAnnotations(t *testing.T) {
	err, annotations := ParseBoostAnnotations(filepath.Join("..", "testdata", "boost", "testlog.xml"))
	if err != nil {
		t.Errorf("Errored: %v", err)
	}
	if len(annotations) != 2 {
		t.Errorf("Wrong amount of annotations: %v", len(annotations))
	}
	expected := model.Annotation{
		Level:      model.LevelFailure,
		Message:    "an error happened",
		RawDetails: "",
		Path:       "test/boost_tests.cpp",
		StartLine:  130,
		EndLine:    130,
	}
	if diff := deep.Equal(expected, annotations[0]); diff != nil {
		t.Error(diff)
	}

	expected = model.Annotation{
		Level:      model.LevelFailure,
		Message:    "another error",
		RawDetails: "",
		Path:       "test/boost_tests2.cpp",
		StartLine:  230,
		EndLine:    230,
	}
	if diff := deep.Equal(expected, annotations[1]); diff != nil {
		t.Error(diff)
	}

}
