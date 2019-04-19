package annotations

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/go-test/deep"
	"path/filepath"
	"testing"
)

func Test_DefaultValue(t *testing.T) {
	err, annotations := ParseAnnotations("junit", filepath.Join("testdata", "junit", "TEST-LibraryTest.xml"))
	if err != nil {
		t.Errorf("Errored: %v", err)
	}
	if len(annotations) != 1 {
		t.Errorf("Wrong amount of annotations: %v", len(annotations))
	}
	annotation := annotations[0]
	annotation.RawDetails = ""
	expected := model.Annotation{
		Type:               model.TestResultAnnotationType,
		Level:              "failure",
		Message:            "LibraryTest.testSomeLibraryMethod",
		FullyQualifiedName: "LibraryTest.testSomeLibraryMethod",
		RawDetails:         "",
		Location: &model.FileLocation{
			Path:      "LibraryTest.java",
			StartLine: 10,
			EndLine:   10,
		},
	}

	if diff := deep.Equal(expected, annotation); diff != nil {
		t.Error(diff)
	}
}
