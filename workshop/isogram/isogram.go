package isogram

import "unicode"

/*
Determine if a word or phrase is an isogram.

An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

Examples of isograms:

lumberjacks
background
downstream
six-year-old
The word isograms, however, is not an isogram, because the s repeats.
*/

// IsIsogram is a function that determines if a word or phrase is an isogram.
func IsIsogram(s string) bool {
	seen := map[rune]bool{}
	for _, char := range s {
		char = unicode.ToLower(char)
		if 'a' <= char && char <= 'z' {
			if _, ok := seen[char]; ok {
				return false
			}
			seen[char] = true
		}
	}
	return true
}
