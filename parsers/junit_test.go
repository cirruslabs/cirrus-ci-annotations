package parsers

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/go-test/deep"
	"path/filepath"
	"testing"
)

func Test_JunitJava(t *testing.T) {
	err, annotations := ParseJUnitAnnotations(filepath.Join("..", "testdata", "junit", "JunitJava.xml"))
	if err != nil {
		t.Errorf("Errored: %v", err)
	}
	if len(annotations) != 1 {
		t.Errorf("Wrong amount of annotations: %v", len(annotations))
	}
	annotation := annotations[0]
	annotation.RawDetails = ""
	expected := model.Annotation{
		Level:      model.LevelFailure,
		Message:    "LibraryTest.testSomeLibraryMethod",
		RawDetails: "",
		Path:       "LibraryTest.java",
		StartLine:  10,
		EndLine:    10,
	}

	if diff := deep.Equal(expected, annotation); diff != nil {
		t.Error(diff)
	}
}

func Test_JunitKotlin(t *testing.T) {
	err, annotations := ParseJUnitAnnotations(filepath.Join("..", "testdata", "junit", "JunitKotlin.xml"))
	if err != nil {
		t.Errorf("Errored: %v", err)
	}
	if len(annotations) != 1 {
		t.Errorf("Wrong amount of annotations: %v", len(annotations))
	}
	annotation := annotations[0]
	annotation.RawDetails = ""
	expected := model.Annotation{
		Level:      model.LevelFailure,
		Message:    "com.fkorotkov.kubernetes.SimpleCompilationTest.testService",
		RawDetails: "",
		Path:       "SimpleCompilationTest.kt",
		StartLine:  41,
		EndLine:    41,
	}

	if diff := deep.Equal(expected, annotation); diff != nil {
		t.Error(diff)
	}
}

func Test_PythonXMLRunner(t *testing.T) {
	err, annotations := ParseJUnitAnnotations(filepath.Join("..", "testdata", "junit", "PythonXMLRunner.xml"))
	if err != nil {
		t.Errorf("Errored: %v", err)
	}
	if len(annotations) != 1 {
		t.Errorf("Wrong amount of annotations: %v", len(annotations))
	}
	annotation := annotations[0]
	annotation.RawDetails = ""
	expected := model.Annotation{
		Level:      model.LevelFailure,
		Message:    "tests.Tests.test_utilities",
		RawDetails: "",
		Path:       "tests.py",
		StartLine:  70,
		EndLine:    70,
	}

	if diff := deep.Equal(expected, annotation); diff != nil {
		t.Error(diff)
	}
}
