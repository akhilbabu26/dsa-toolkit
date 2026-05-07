package hashmap

// Set is a generic-style string set backed by map[string]struct{}.
//
// # Why map[string]struct{} instead of map[string]bool?
//
// struct{} is an empty struct — it occupies exactly 0 bytes of memory.
// map[string]bool uses 1 byte per entry for the bool value.
// For large sets, struct{} is the idiomatic Go choice for pure membership tracking.
//
//	s := NewSet()
//	s.Add("apple")
//	s.Contains("apple")   → true
//	s.Contains("banana")  → false
//	s.Remove("apple")
//	s.Size()              → 0
type Set struct {
	data map[string]struct{}
}

// NewSet creates and returns a new empty Set.
func NewSet() *Set {
	return &Set{data: make(map[string]struct{})}
}

// Add inserts an element into the set.
// If the element already exists, this is a no-op (sets contain no duplicates).
//
// Time Complexity: O(1) average
func (s *Set) Add(item string) {
	// struct{}{} is a zero-size struct literal — the value we store in the map.
	s.data[item] = struct{}{}
}

// Remove deletes an element from the set.
// If the element is not present, this is a no-op (Go's delete is safe on missing keys).
//
// Time Complexity: O(1) average
func (s *Set) Remove(item string) {
	delete(s.data, item)
}

// Contains reports whether item is in the set.
//
// The two-value map lookup (`_, ok := m[key]`) is the correct way to
// test for key existence in Go — using just `m[key]` would return
// the zero value (struct{}{}) even if the key is absent, making it
// indistinguishable from a real entry.
//
// Time Complexity: O(1) average
func (s *Set) Contains(item string) bool {
	_, ok := s.data[item]
	return ok
}

// Size returns the number of elements currently in the set.
//
// Time Complexity: O(1) — Go's len() on a map is O(1)
func (s *Set) Size() int {
	return len(s.data)
}

// Elements returns all items in the set as an unordered slice.
// Note: map iteration order in Go is intentionally randomised — do not
// rely on the order of the returned slice.
//
// Time Complexity: O(n)
func (s *Set) Elements() []string {
	items := make([]string, 0, len(s.data))
	for k := range s.data {
		items = append(items, k)
	}
	return items
}

// Union returns a new Set containing all elements from both sets.
//
// Time Complexity: O(n + m)
func Union(a, b *Set) *Set {
	result := NewSet()
	for k := range a.data {
		result.Add(k)
	}
	for k := range b.data {
		result.Add(k)
	}
	return result
}

// Intersection returns a new Set with only elements present in both sets.
// We iterate over the smaller set for efficiency.
//
// Time Complexity: O(min(n, m))
func Intersection(a, b *Set) *Set {
	result := NewSet()
	// Iterate over whichever set is smaller
	small, large := a, b
	if a.Size() > b.Size() {
		small, large = b, a
	}
	for k := range small.data {
		if large.Contains(k) {
			result.Add(k)
		}
	}
	return result
}
