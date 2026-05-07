package hashmap

import (
	"reflect"
	"sort"
	"testing"
)

// ── WordFrequency Tests ───────────────────────────────────────────────────────

func TestWordFrequency(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{
			name:     "basic frequency",
			input:    []string{"go", "is", "great", "go", "is", "fun"},
			expected: map[string]int{"go": 2, "is": 2, "great": 1, "fun": 1},
		},
		{
			name:     "case normalisation",
			input:    []string{"Go", "go", "GO"},
			expected: map[string]int{"go": 3},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: map[string]int{},
		},
		{
			name:     "single word",
			input:    []string{"hello"},
			expected: map[string]int{"hello": 1},
		},
		{
			name:     "blank entries skipped",
			input:    []string{"go", "", "  ", "go"},
			expected: map[string]int{"go": 2},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := WordFrequency(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("WordFrequency(%v)\n  got:  %v\n  want: %v", tc.input, got, tc.expected)
			}
		})
	}
}

// ── FindDuplicates Tests ──────────────────────────────────────────────────────

func TestFindDuplicates(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "has duplicates",
			input:    []string{"apple", "banana", "apple", "cherry", "banana"},
			expected: []string{"apple", "banana"},
		},
		{
			name:     "case insensitive duplicates",
			input:    []string{"Go", "go", "Python"},
			expected: []string{"go"},
		},
		{
			name:     "no duplicates",
			input:    []string{"a", "b", "c"},
			expected: nil,
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindDuplicates(tc.input)
			// Sort both slices for deterministic comparison (map iteration is random)
			sort.Strings(got)
			sort.Strings(tc.expected)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("FindDuplicates(%v)\n  got:  %v\n  want: %v", tc.input, got, tc.expected)
			}
		})
	}
}

// ── Set Tests ─────────────────────────────────────────────────────────────────

func TestSet(t *testing.T) {
	s := NewSet()

	if s.Size() != 0 {
		t.Errorf("new set should be empty, got size %d", s.Size())
	}

	s.Add("go")
	s.Add("python")
	s.Add("go") // duplicate — should be ignored

	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}

	if !s.Contains("go") {
		t.Error("expected set to contain 'go'")
	}

	if s.Contains("ruby") {
		t.Error("expected set NOT to contain 'ruby'")
	}

	s.Remove("go")
	if s.Contains("go") {
		t.Error("expected 'go' to be removed")
	}

	// Remove non-existent element — must not panic
	s.Remove("nonexistent")
}

func TestSetUnion(t *testing.T) {
	a := NewSet()
	a.Add("go")
	a.Add("python")

	b := NewSet()
	b.Add("python")
	b.Add("rust")

	u := Union(a, b)
	for _, item := range []string{"go", "python", "rust"} {
		if !u.Contains(item) {
			t.Errorf("union should contain %q", item)
		}
	}
	if u.Size() != 3 {
		t.Errorf("union size: got %d, want 3", u.Size())
	}
}

func TestSetIntersection(t *testing.T) {
	a := NewSet()
	a.Add("go")
	a.Add("python")

	b := NewSet()
	b.Add("python")
	b.Add("rust")

	i := Intersection(a, b)
	if !i.Contains("python") {
		t.Error("intersection should contain 'python'")
	}
	if i.Contains("go") || i.Contains("rust") {
		t.Error("intersection should only contain 'python'")
	}
}
