package vowels

import "strings"

var vowelsToNumber = map[string]string{
	"a": "1",
	"e": "2",
	"i": "3",
	"o": "4",
	"u": "5",
}

var numberToVowels = map[string]string{
	"1": "a",
	"2": "e",
	"3": "i",
	"4": "o",
	"5": "u",
}

func Encode(word string) string {
	var result []string
	for _, l := range word {
		char := string(l)
		replace, ok := vowelsToNumber[strings.ToLower(char)]
		if !ok {
			result = append(result, char)
		} else {
			result = append(result, replace)
		}
	}
	return strings.Join(result, "")
}

func Decode(word string) string {
	var result []string
	for _, l := range word {
		char := string(l)
		replace, ok := numberToVowels[char]
		if !ok {
			result = append(result, char)
		} else {
			result = append(result, replace)
		}
	}
	return strings.Join(result, "")
}
