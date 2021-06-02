package parsers

import (
	"path/filepath"
	"testing"

	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/go-test/deep"
)

func Test_AndroidLint_Multiple_Locations(t *testing.T) {
	err, annotations := ParseAndroidLintAnnotations(filepath.Join("..", "testdata", "android-lint", "lint-results-multiple-locations.xml"))
	if err != nil {
		t.Errorf("Errored: %v", err)
	}
	if len(annotations) != 2 {
		t.Errorf("Wrong amount of annotations: %v", len(annotations))
	}

	firstAnnotation := annotations[0]
	firstExpected := model.Annotation{
		Level:       model.LevelWarning,
		Message:     "The resource `R.string.my_string` appears to be unused",
		RawDetails:  "Unused resources make applications larger and slow down builds.\n\nThe unused resource check can ignore tests. If you want to include resources that are only referenced from tests, consider packaging them in a test source set instead.\n\nYou can include test sources in the unused resource check by setting the system property lint.unused-resources.include-tests=true, and to exclude them (usually for performance reasons), use lint.unused-resources.exclude-tests=true.",
		Path:        "/path/to/project/app/src/main/res/values-de/strings.xml",
		StartLine:   35,
		EndLine:     35,
		StartColumn: 11,
		EndColumn:   11,
	}

	if diff := deep.Equal(firstExpected, firstAnnotation); diff != nil {
		t.Error(diff)
	}

	secondAnnotation := annotations[1]
	secondExpected := model.Annotation{
		Level:       model.LevelWarning,
		Message:     "The resource `R.string.my_string` appears to be unused",
		RawDetails:  "Unused resources make applications larger and slow down builds.\n\nThe unused resource check can ignore tests. If you want to include resources that are only referenced from tests, consider packaging them in a test source set instead.\n\nYou can include test sources in the unused resource check by setting the system property lint.unused-resources.include-tests=true, and to exclude them (usually for performance reasons), use lint.unused-resources.exclude-tests=true.",
		Path:        "/path/to/project/app/src/main/res/values-fr/strings.xml",
		StartLine:   37,
		EndLine:     37,
		StartColumn: 11,
		EndColumn:   11,
	}

	if diff := deep.Equal(secondExpected, secondAnnotation); diff != nil {
		t.Error(diff)
	}
}

func Test_AndroidLint_Multiple_Issues(t *testing.T) {
	err, annotations := ParseAndroidLintAnnotations(filepath.Join("..", "testdata", "android-lint", "lint-results-multiple-issues.xml"))
	if err != nil {
		t.Errorf("Errored: %v", err)
	}
	if len(annotations) != 3 {
		t.Errorf("Wrong amount of annotations: %v", len(annotations))
	}

	firstAnnotation := annotations[0]
	firstExpected := model.Annotation{
		Level:       model.LevelFailure,
		Message:     "The resource `R.string.my_string` appears to be unused",
		RawDetails:  "Unused resources make applications larger and slow down builds.\n\nThe unused resource check can ignore tests. If you want to include resources that are only referenced from tests, consider packaging them in a test source set instead.\n\nYou can include test sources in the unused resource check by setting the system property lint.unused-resources.include-tests=true, and to exclude them (usually for performance reasons), use lint.unused-resources.exclude-tests=true.",
		Path:        "/path/to/project/app/src/main/res/values-de/strings.xml",
		StartLine:   35,
		EndLine:     35,
		StartColumn: 11,
		EndColumn:   11,
	}

	if diff := deep.Equal(firstExpected, firstAnnotation); diff != nil {
		t.Error(diff)
	}

	secondAnnotation := annotations[1]
	secondExpected := model.Annotation{
		Level:       model.LevelFailure,
		Message:     "The resource `R.string.my_string` appears to be unused",
		RawDetails:  "Unused resources make applications larger and slow down builds.\n\nThe unused resource check can ignore tests. If you want to include resources that are only referenced from tests, consider packaging them in a test source set instead.\n\nYou can include test sources in the unused resource check by setting the system property lint.unused-resources.include-tests=true, and to exclude them (usually for performance reasons), use lint.unused-resources.exclude-tests=true.",
		Path:        "/path/to/project/app/src/main/res/values-de/strings.xml",
		StartLine:   35,
		EndLine:     35,
		StartColumn: 11,
		EndColumn:   11,
	}

	if diff := deep.Equal(secondExpected, secondAnnotation); diff != nil {
		t.Error(diff)
	}

	thirdAnnotation := annotations[2]
	thirdExpected := model.Annotation{
		Level:       model.LevelNotice,
		Message:     "The resource `R.string.my_string` appears to be unused",
		RawDetails:  "Unused resources make applications larger and slow down builds.\n\nThe unused resource check can ignore tests. If you want to include resources that are only referenced from tests, consider packaging them in a test source set instead.\n\nYou can include test sources in the unused resource check by setting the system property lint.unused-resources.include-tests=true, and to exclude them (usually for performance reasons), use lint.unused-resources.exclude-tests=true.",
		Path:        "/path/to/project/app/src/main/res/values-de/strings.xml",
		StartLine:   35,
		EndLine:     35,
		StartColumn: 11,
		EndColumn:   11,
	}

	if diff := deep.Equal(thirdExpected, thirdAnnotation); diff != nil {
		t.Error(diff)
	}
}

func Test_AndroidLint_Ignored_Issue(t *testing.T) {
	err, annotations := ParseAndroidLintAnnotations(filepath.Join("..", "testdata", "android-lint", "lint-results-ignored-issue.xml"))
	if err != nil {
		t.Errorf("Errored: %v", err)
	}
	if len(annotations) != 0 {
		t.Errorf("Wrong amount of annotations: %v", len(annotations))
	}
}
