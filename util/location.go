package util

import (
	"regexp"
	"strconv"
	"strings"
)

func GuessLocationIgnored(data string, ignorePatters []string) (string, int64, int64) {
	regex, err := regexp.Compile(`([\w\\.]+):(\d+)`)
	if err != nil {
		return "", 0, 0
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
		return parts[0], lineNumber, lineNumber
	}
	return "", 0, 0
}
