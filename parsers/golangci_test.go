package parsers

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/go-test/deep"
	"path/filepath"
	"testing"
)

func Test_GoLangCI(t *testing.T) {
	err, annotations := ParseGoLangCIAnnotations(filepath.Join("..", "testdata", "json", "golangci.json"))
	if err != nil {
		t.Errorf("Errored: %v", err)
	}
	if len(annotations) != 1 {
		t.Errorf("Wrong amount of annotations: %v", len(annotations))
	}
	annotation := annotations[0]
	annotation.RawDetails = ""
	expected := &model.Annotation{
		Level:       model.LevelFailure,
		Message:     "S1007: should use raw string (`...`) with regexp.Compile to avoid having to escape twice (gosimple)",
		RawDetails:  "",
		Path:        "util/location.go",
		StartLine:   11,
		EndLine:     11,
		StartColumn: 16,
	}

	if diff := deep.Equal(expected, annotation); diff != nil {
		t.Error(diff)
	}
}
