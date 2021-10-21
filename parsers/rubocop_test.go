package parsers

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/go-test/deep"
	"path/filepath"
	"testing"
)

// Test_RuboCop_DocsExample ensures that JSON example from the documentation[1] is parsed properly.
// [1]: https://github.com/rubocop-hq/rubocop/blob/master/docs/modules/ROOT/pages/formatters.adoc#json-formatter
func Test_RuboCop_DocsExample(t *testing.T) {
	var expectedAnnotations = []model.Annotation{
		{
			Level:       model.LevelNotice,
			Message:     "LineLength: Line is too long. [81/80]",
			Path:        "lib/bar.rb",
			StartLine:   546,
			EndLine:     546,
			StartColumn: 80,
			EndColumn:   80,
		},
		{
			Level:       model.LevelWarning,
			Message:     "UnreachableCode: Unreachable code detected.",
			Path:        "lib/bar.rb",
			StartLine:   15,
			EndLine:     15,
			StartColumn: 9,
			EndColumn:   9,
		},
	}

	err, gotAnnotations := ParseRuboCopAnnotations(filepath.Join("..", "testdata", "json", "rubocop.json"))
	if err != nil {
		t.Error(err)
		return
	}

	if diff := deep.Equal(expectedAnnotations, gotAnnotations); diff != nil {
		t.Error(diff)
	}
}
