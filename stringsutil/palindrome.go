package stringsutil

import (
	"strings"
	"unicode"
)

// PalindromeResult holds both the verdict and the cleaned string used for
// comparison, making it useful for display in a CLI or test output.
type PalindromeResult struct {
	Original string // the string as passed in
	Cleaned  string // normalized form used for comparison
	IsPalin  bool   // true if the cleaned string is a palindrome
}

// IsPalindrome checks whether s is a palindrome.
//
// Normalization rules (beginner-friendly defaults):
//  1. Convert to lowercase — "Racecar" == "racecar"
//  2. Strip non-alphanumeric characters — "A man, a plan, a canal: Panama" → true
//
// This approach uses rune-safe operations so it works correctly with
// Unicode input (e.g., accented characters).
//
// Examples:
//
//	"racecar"                        → true
//	"A man, a plan, a canal: Panama" → true
//	"hello"                          → false
//	""                               → true  (empty string is vacuously a palindrome)
func IsPalindrome(s string) PalindromeResult {
	// Step 1: lowercase
	lower := strings.ToLower(s)

	// Step 2: keep only letters and digits (unicode-aware)
	var cleaned []rune
	for _, r := range lower {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}

	// Step 3: two-pointer palindrome check on the rune slice
	n := len(cleaned)
	isPalin := true
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			isPalin = false
			break
		}
	}

	return PalindromeResult{
		Original: s,
		Cleaned:  string(cleaned),
		IsPalin:  isPalin,
	}
}
