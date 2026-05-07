// Package stringsutil provides UTF-8-safe string utility functions.
//
// # Why runes matter in Go
//
// A Go string is a sequence of bytes, not characters. ASCII characters are
// 1 byte each, but many Unicode characters (e.g., emojis, accented letters,
// CJK ideographs) span 2–4 bytes. Reversing a string byte-by-byte would
// corrupt multi-byte sequences.
//
// The fix: convert to []rune first. A rune (alias for int32) holds a single
// Unicode code point, so reversing []rune is always safe.
//
//	"café"  →  bytes: [99 97 102 195 169]  (5 bytes, 4 runes)
//	Correct reverse: "éfac"
//	Byte-level reverse would corrupt the 'é' (two-byte sequence 195 169).
package stringsutil

// Reverse returns the Unicode-correct reversal of the given string.
//
// Example:
//
//	Reverse("hello")  → "olleh"
//	Reverse("café")   → "éfac"
//	Reverse("🙂😎")    → "😎🙂"
//	Reverse("")        → ""
func Reverse(s string) string {
	// Convert to a slice of runes so each element is one full Unicode character.
	runes := []rune(s)

	// Two-pointer swap: left pointer moves right, right pointer moves left,
	// until they meet in the middle.
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	// Convert back to a string. Go's string() conversion from []rune
	// correctly encodes each rune back into its UTF-8 byte sequence.
	return string(runes)
}
