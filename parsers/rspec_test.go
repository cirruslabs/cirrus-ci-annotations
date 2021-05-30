package parsers

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"github.com/go-test/deep"
	"path/filepath"
	"testing"
)

// Test_RSpec_MultipleStates ensures that "pending" and "failed" RSpec example states are handled properly,
// while the "passed" state is skipped.
//
// The following dummy_spec.rb was used to generate the test data:
//
// class Dummy
// end
//
// RSpec.describe Dummy do
//   it "passes with deprecation warning" do
//     expect {
//       nil
//     }.not_to raise_error(ArgumentError)
//   end
//
//   it "gets skipped" do
//     pending("not implemented yet")
//     expect(1).to eq(2)
//   end
//
//   it "fails" do
//     expect(1).to eq(2)
//   end
// end
func Test_RSpec_MultipleStates(t *testing.T) {
	var expectedAnnotations = []model.Annotation{
		{
			Level:              model.LevelNotice,
			Message:            "Dummy gets skipped",
			RawDetails:         "not implemented yet",
			Location: &model.FileLocation{
				Path:      "spec/dummy_spec.rb",
				StartLine: 11,
				EndLine:   11,
			},
		},
		{
			Level:              model.LevelFailure,
			Message:            "Dummy fails",
			RawDetails:         "\nexpected: 2\n     got: 1\n\n(compared using ==)\n",
			Location: &model.FileLocation{
				Path:      "spec/dummy_spec.rb",
				StartLine: 16,
				EndLine:   16,
			},
		},
	}

	err, gotAnnotations := ParseRSpecAnnotations(filepath.Join("..", "testdata", "json", "rspec.json"))
	if err != nil {
		t.Error(err)
		return
	}

	if diff := deep.Equal(expectedAnnotations, gotAnnotations); diff != nil {
		t.Error(diff)
	}
}
