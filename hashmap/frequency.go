// Package hashmap demonstrates idiomatic Go map usage for two common DSA
// patterns: frequency counting and set operations.
//
// # How Go maps work (beginner primer)
//
// A Go map is a hash table — it maps keys to values in O(1) average time.
//
//	m := make(map[string]int)
//	m["apple"]++      // zero value for int is 0, so this works on first access
//	count := m["banana"]  // returns 0 (zero value) if key doesn't exist
//
// The two-value form lets you distinguish "key missing" from "key has zero value":
//
//	val, ok := m["key"]
//	if !ok { /* key is not in the map */ }
package hashmap

import "strings"

// WordFrequency takes a slice of words and returns a map of how often each
// word appears. All words are normalised to lowercase before counting so that
// "Go" and "go" are treated as the same word.
//
// Time Complexity:  O(n) — single pass over the slice
// Space Complexity: O(k) — where k is the number of unique words
//
// Example:
//
//	WordFrequency([]string{"go", "is", "great", "go", "is", "fun"})
//	→ map[go:2 is:2 great:1 fun:1]
func WordFrequency(words []string) map[string]int {
	freq := make(map[string]int, len(words)) // pre-size hint for efficiency

	for _, w := range words {
		// Normalise: trim whitespace and lowercase so "Go", "go", "GO" all count
		key := strings.ToLower(strings.TrimSpace(w))
		if key == "" {
			continue // skip blank entries
		}
		freq[key]++ // Go's zero-value means this works even on first encounter
	}

	return freq
}

// CharFrequency counts the occurrences of each rune in the given string.
// It uses rune iteration (not byte) so it handles UTF-8 correctly.
//
// Example:
//
//	CharFrequency("hello") → map[h:1 e:1 l:2 o:1]
func CharFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, r := range s { // ranging over a string yields runes, not bytes
		freq[r]++
	}
	return freq
}

// FindDuplicates returns words that appear more than once in the input slice.
// Words are case-normalised before comparison.
//
// Example:
//
//	FindDuplicates([]string{"apple", "banana", "Apple", "cherry", "banana"})
//	→ ["apple", "banana"]
func FindDuplicates(words []string) []string {
	freq := WordFrequency(words) // reuse WordFrequency — no logic duplication

	var duplicates []string
	// Use a secondary set to avoid adding the same duplicate twice.
	seen := make(map[string]bool)
	for _, w := range words {
		key := strings.ToLower(strings.TrimSpace(w))
		if freq[key] > 1 && !seen[key] {
			duplicates = append(duplicates, key)
			seen[key] = true
		}
	}
	return duplicates
}
