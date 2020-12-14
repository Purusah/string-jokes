package latin

import (
	"strings"
	"unicode"
)

const suffix = "ay"
const vowels = "aeiou"

func Encode(word string) string {
	var sb strings.Builder
	for ix, ch := range word {
		char := string(ch)
		if ix == 0 && unicode.IsDigit(ch) {
			return word
		}
		if strings.Contains(vowels, strings.ToLower(char)) {
			sb.WriteString(word[ix:])
			sb.WriteString(word[:ix])
			break
		}
	}
	sb.WriteString(suffix)
	return sb.String()
}
