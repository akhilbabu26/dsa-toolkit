package stringsutil

import (
	"testing"
)

// ── Reverse Tests ─────────────────────────────────────────────────────────────

func TestReverse(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple ascii", "hello", "olleh"},
		{"single char", "a", "a"},
		{"empty string", "", ""},
		{"palindrome unchanged", "racecar", "racecar"},
		{"with spaces", "hello world", "dlrow olleh"},
		// UTF-8 safety — multi-byte characters must not be corrupted
		{"utf8 accented", "café", "éfac"},
		{"utf8 emoji", "🙂😎", "😎🙂"},
		{"utf8 cjk", "你好", "好你"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Reverse(tc.input)
			if got != tc.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tc.input, got, tc.expected)
			}
		})
	}
}

// ── Palindrome Tests ──────────────────────────────────────────────────────────

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"simple palindrome", "racecar", true},
		{"simple non-palindrome", "hello", false},
		{"empty string", "", true},
		{"single char", "a", true},
		// Normalization: case + punctuation should be ignored
		{"mixed case", "RaceCar", true},
		{"with spaces and punctuation", "A man, a plan, a canal: Panama", true},
		{"with spaces only", "Was it a car or a cat I saw?", true},
		{"numbers palindrome", "12321", true},
		{"numbers non-palindrome", "12345", false},
		{"all same chars", "aaaa", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsPalindrome(tc.input)
			if result.IsPalin != tc.expected {
				t.Errorf("IsPalindrome(%q).IsPalin = %v; want %v (cleaned: %q)",
					tc.input, result.IsPalin, tc.expected, result.Cleaned)
			}
		})
	}
}

// ── VowelCounter Tests ────────────────────────────────────────────────────────

func TestCountVowels(t *testing.T) {
	cases := []struct {
		name          string
		input         string
		expectedCount int
	}{
		{"normal sentence", "Hello, World!", 3},
		{"all vowels", "aeiouAEIOU", 10},
		{"no vowels", "rhythm", 0},
		{"empty string", "", 0},
		{"single vowel", "A", 1},
		{"single consonant", "b", 0},
		// Non-ASCII chars that look like vowels should not be counted.
		// "café" = c + a + f + é (é is U+00E9, a 2-byte rune — NOT ASCII 'e')
		// So only 'a' is a counted vowel → expected count is 1.
		{"accented vowel not counted", "café", 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := CountVowels(tc.input)
			if result.Count != tc.expectedCount {
				t.Errorf("CountVowels(%q).Count = %d; want %d", tc.input, result.Count, tc.expectedCount)
			}
		})
	}
}

// TestCountVowelsFrequency verifies the per-vowel frequency breakdown.
func TestCountVowelsFrequency(t *testing.T) {
	result := CountVowels("Hello Beautiful")
	// H-e-l-l-o  B-e-a-u-t-i-f-u-l
	// vowels: e, o, e, a, u, i, u → a:1, e:2, i:1, o:1, u:2
	wantFreq := map[rune]int{'e': 2, 'o': 1, 'a': 1, 'u': 2, 'i': 1}
	for vowel, want := range wantFreq {
		if got := result.Frequency[vowel]; got != want {
			t.Errorf("Frequency[%q] = %d; want %d", vowel, got, want)
		}
	}
}
