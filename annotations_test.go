package annotations

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/go-test/deep"
	"path/filepath"
	"testing"
)

func Test_JunitJava(t *testing.T) {
	err, annotations := ParseAnnotations("junit", filepath.Join("testdata", "junit", "JunitJava.xml"))
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

func Test_JunitKotlin(t *testing.T) {
	err, annotations := ParseAnnotations("junit", filepath.Join("testdata", "junit", "JunitKotlin.xml"))
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
		Message:            "com.fkorotkov.kubernetes.SimpleCompilationTest.testService",
		FullyQualifiedName: "com.fkorotkov.kubernetes.SimpleCompilationTest.testService",
		RawDetails:         "",
		Location: &model.FileLocation{
			Path:      "SimpleCompilationTest.kt",
			StartLine: 41,
			EndLine:   41,
		},
	}

	if diff := deep.Equal(expected, annotation); diff != nil {
		t.Error(diff)
	}
}

func Test_PythonXMLRunner(t *testing.T) {
	err, annotations := ParseAnnotations("junit", filepath.Join("testdata", "junit", "PythonXMLRunner.xml"))
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
		Message:            "tests.Tests.test_utilities",
		FullyQualifiedName: "tests.Tests.test_utilities",
		RawDetails:         "",
		Location: &model.FileLocation{
			Path:      "tests.py",
			StartLine: 70,
			EndLine:   70,
		},
	}

	if diff := deep.Equal(expected, annotation); diff != nil {
		t.Error(diff)
	}
}

func Test_GoJUnitReport(t *testing.T) {
	err, annotations := ParseAnnotations("junit", filepath.Join("testdata", "junit", "GoJUnitReport.xml"))
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
		Message:            "cirrus-ci-annotations.Test_PythonXMLRunner",
		FullyQualifiedName: "cirrus-ci-annotations.Test_PythonXMLRunner",
		RawDetails:         "",
		Location: &model.FileLocation{
			Path:      "annotations_test.go",
			StartLine: 90,
			EndLine:   90,
		},
	}

	if diff := deep.Equal(expected, annotation); diff != nil {
		t.Error(diff)
	}
}
