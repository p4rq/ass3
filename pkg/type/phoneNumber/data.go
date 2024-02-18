package phoneNumber

import (
	"strings"
	"unicode"
)

func getNumbers(input string) string {
	var builder strings.Builder

	for _, t := range input {
		if unicode.IsDigit(t) {
			builder.WriteRune(t)
		}
	}

	return builder.String()
}
