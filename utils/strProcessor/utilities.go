package utils

import (
    "strings"
)

func splitString (input string) []string{
    return strings.Split(input, ", ")
}
