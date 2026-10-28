package repl

import (
	"strings"
)

func CleanInput(text string) []string {
	var result []string
	trimmedStr := strings.TrimSpace(text)
	loweredStr := strings.ToLower(trimmedStr)
	splitStr := strings.Split(loweredStr, " ")

	for _, str := range splitStr {
		finalTrimmedString := strings.TrimSpace(str)
		if finalTrimmedString != "" {
			result = append(result, finalTrimmedString)
		}
	}

	return result
}
