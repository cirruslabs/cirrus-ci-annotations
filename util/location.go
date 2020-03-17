package util

import (
	"github.com/cirruslabs/cirrus-ci-annotations/model"
	"regexp"
	"strconv"
	"strings"
)

func GuessLocationIgnored(data string, ignorePatters []string) *model.FileLocation {
	regex, err := regexp.Compile("([\\w\\.]+)\\:(\\d+)")
	if err != nil {
		return nil
	}
	lines := strings.Split(data, "\n")
LinesLoop:
	for _, line := range lines {
		path := regex.FindString(line)
		if path == "" {
			continue
		}
		for _, ignorePatter := range ignorePatters {
			if strings.Contains(line, ignorePatter) {
				continue LinesLoop
			}
		}
		parts := strings.Split(path, ":")
		lineNumber, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			continue
		}
		return &model.FileLocation{
			Path:      parts[0],
			StartLine: lineNumber,
			EndLine:   lineNumber,
		}
	}
	return nil
}
