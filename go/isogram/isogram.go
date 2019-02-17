package isogram

import (
	"unicode"
)

// IsIsogram determines whether a string is an isogram or not
func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}

		letter = unicode.ToLower(letter)
		if letters[letter] {
			return false
		}
		letters[letter] = true
	}
	return true
}