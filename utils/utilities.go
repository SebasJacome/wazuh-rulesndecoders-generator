package utils

import (
	"strings"
)

func SplitString(input string) []string {
	return strings.Split(input, ",")
}

func CompareLogAndVars(log string, vars []string) bool {
	for _, value := range vars {
		if !strings.Contains(log, value) {
			return false
		}
	}
	return true
}
