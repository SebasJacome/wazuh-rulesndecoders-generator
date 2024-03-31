package utils

import (
	"strings"
)

func SplitString(input string) []string {
	return strings.Split(input, ",")
}

func CompareLogAndVars(log string, vars []string) string {
	var result string
	for _, value := range vars {
		if !strings.Contains(log, value) {
			result += value + ","
		}
	}

	if result != "" {
		result = result[:len(result)-1]
	}
	return result
}

func CompareExistingIDs(ruleIDs map[string]bool, ID string) bool {
	if ruleIDs[ID] {
		return true
	}
	return false
}
