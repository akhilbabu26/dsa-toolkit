package sorting

import (
	"reflect"
	"testing"
)

// TestCase holds a single test scenario used across all sorting algorithms.
type TestCase struct {
	name     string
	input    []int
	expected []int
}

// sharedCases lists test scenarios that every sorting algorithm must pass.
// Table-driven tests like this are idiomatic Go — they avoid repetition and
// make it trivial to add new edge cases without touching algorithm logic.
var sharedCases = []TestCase{
	{
		name:     "unsorted slice",
		input:    []int{5, 3, 8, 1, 9, 2},
		expected: []int{1, 2, 3, 5, 8, 9},
	},
	{
		name:     "already sorted",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "reverse sorted",
		input:    []int{9, 7, 5, 3, 1},
		expected: []int{1, 3, 5, 7, 9},
	},
	{
		name:     "single element",
		input:    []int{42},
		expected: []int{42},
	},
	{
		name:     "empty slice",
		input:    []int{},
		expected: []int{},
	},
	{
		name:     "duplicate values",
		input:    []int{4, 2, 4, 1, 2},
		expected: []int{1, 2, 2, 4, 4},
	},
	{
		name:     "negative numbers",
		input:    []int{-3, 0, -7, 5, -1},
		expected: []int{-7, -3, -1, 0, 5},
	},
	{
		name:     "all same",
		input:    []int{3, 3, 3, 3},
		expected: []int{3, 3, 3, 3},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, tc := range sharedCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BubbleSort(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("BubbleSort(%v)\n  got:  %v\n  want: %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tc := range sharedCases {
		t.Run(tc.name, func(t *testing.T) {
			got := InsertionSort(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("InsertionSort(%v)\n  got:  %v\n  want: %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tc := range sharedCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SelectionSort(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("SelectionSort(%v)\n  got:  %v\n  want: %v", tc.input, got, tc.expected)
			}
		})
	}
}

// TestSortDoesNotMutateInput verifies that all algorithms work on copies —
// the original slice passed by the caller must remain unchanged.
func TestSortDoesNotMutateInput(t *testing.T) {
	original := []int{5, 3, 8, 1}
	snapshot := []int{5, 3, 8, 1}

	BubbleSort(original)
	InsertionSort(original)
	SelectionSort(original)

	if !reflect.DeepEqual(original, snapshot) {
		t.Errorf("sorting functions must not mutate input: got %v, want %v", original, snapshot)
	}
}
