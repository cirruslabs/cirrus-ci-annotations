package annotations_test

import (
	"github.com/cirruslabs/cirrus-ci-annotations"
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateAnnotations(t *testing.T) {
	rawAnnotations := []model.Annotation{
		{
			Path: "JunitJava.xml",
		},
	}
	expectedAnnotations := []model.Annotation{
		{
			Path: "junit/JunitJava.xml",
		},
	}

	processedAnnotations, err := annotations.NormalizeAnnotations("testdata", rawAnnotations)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedAnnotations, processedAnnotations)
}
