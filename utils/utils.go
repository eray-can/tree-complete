package utils

import (
	"strings"
	"unicode"
)

func ToLowerByLang(s string, language unicode.SpecialCase) string {
	var result strings.Builder
	for _, r := range s {
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}
