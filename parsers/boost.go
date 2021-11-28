package parsers

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/cirruslabs/cirrus-ci-annotations/model"
)

type BoostTestError struct {
	XMLName xml.Name `xml:"Error"`
	File    string   `xml:"file,attr"`
	Line    int64    `xml:"line,attr"`
	Message string   `xml:",chardata"`
}

type BoostTestCase struct {
	XMLName xml.Name         `xml:"TestCase"`
	Name    string           `xml:"name,attr"`
	File    string           `xml:"file,attr"`
	Line    int64            `xml:"line,attr"`
	Errors  []BoostTestError `xml:"Error"`
}

type BoostTestSuite struct {
	XMLName   xml.Name         `xml:"TestSuite"`
	Suites    []BoostTestSuite `xml:"TestSuite"`
	TestCases []BoostTestCase  `xml:"TestCase"`
}

func findTestCases(suite *BoostTestSuite) []model.Annotation {
	var annotations []model.Annotation
	for _, s := range suite.Suites {
		annotations = append(annotations, findTestCases(&s)...)
	}
	for _, testCase := range suite.TestCases {
		for _, testError := range testCase.Errors {
			annotations = append(annotations, model.Annotation{
				Level:     model.LevelFailure,
				Message:   testError.Message,
				Path:      testError.File,
				StartLine: testError.Line,
				EndLine:   testError.Line,
			})
		}
	}
	return annotations
}

func ParseBoostAnnotations(path string) (error, []model.Annotation) {
	type suites struct {
		XMLName xml.Name         `xml:"TestLog"`
		Suites  []BoostTestSuite `xml:"TestSuite"`
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}
	results := &suites{}
	if err = xml.Unmarshal(b, results); err != nil {
		return err, nil
	}
	var annotations []model.Annotation
	for _, suite := range results.Suites {
		annotations = append(annotations, findTestCases(&suite)...)
	}
	return nil, annotations
}
