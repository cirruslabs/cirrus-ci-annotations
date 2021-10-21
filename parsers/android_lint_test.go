package parsers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestFormat6(t *testing.T) {
	expected := []model.Annotation{
		{
			Level:       model.LevelWarning,
			Message:     "A newer version of com.android.tools.build:gradle than 3.5.0 is available: 7.0.3. (There is also a newer version of 3.5.𝑥 available, if upgrading to 7.0.3 is difficult: 3.5.4)",
			RawDetails:  "This detector looks for usage of the Android Gradle Plugin where the version you are using is not the current stable release. Using older versions is fine, and there are cases where you deliberately want to stick with an older version. However, you may simply not be aware that a more recent version is available, and that is what this lint check helps find.",
			Path:        "/tmp/cirrus-ci-build/packages/connectivity/connectivity/android/build.gradle",
			StartLine:   12,
			EndLine:     12,
			StartColumn: 9,
			EndColumn:   9,
		},
	}

	err, actual := ParseAndroidLintAnnotations(filepath.Join("..", "testdata", "android-lint", "lint-format-6.xml"))
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}
