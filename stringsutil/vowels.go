package stringsutil

import "strings"

// VowelResult holds detailed statistics about vowels found in a string.
type VowelResult struct {
	Input     string         // original input
	Count     int            // total vowel count
	Vowels    []rune         // each vowel found in order
	Frequency map[rune]int   // per-vowel frequency map
}

// englishVowels is the set of lowercase ASCII vowels.
// We use a map for O(1) lookup instead of scanning a string each time.
var englishVowels = map[rune]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
}

// CountVowels counts the vowels in s (case-insensitive, ASCII vowels only).
//
// The function iterates over runes (not bytes) to handle UTF-8 strings safely.
// Non-ASCII characters that look like vowels (e.g., 'é', 'ü') are NOT counted
// because English vowel rules apply only to the 5 base ASCII vowels.
//
// Examples:
//
//	CountVowels("Hello, World!") → Count: 3  (e, o, o)
//	CountVowels("rhythm")        → Count: 0
//	CountVowels("")              → Count: 0
func CountVowels(s string) VowelResult {
	result := VowelResult{
		Input:     s,
		Frequency: make(map[rune]int),
	}

	// strings.ToLower handles the full Unicode lowercase mapping, but we then
	// only check against the 5 ASCII vowels — safe and predictable.
	for _, r := range strings.ToLower(s) {
		if englishVowels[r] {
			result.Count++
			result.Vowels = append(result.Vowels, r)
			result.Frequency[r]++
		}
	}

	return result
}
