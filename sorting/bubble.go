// Package sorting provides classic sorting algorithm implementations.
// Each algorithm operates on a copy of the input slice to avoid mutating
// the original — a safe, functional-style approach preferred in Go.
package sorting

// BubbleSort repeatedly compares adjacent elements and swaps them if they are
// in the wrong order. The largest unsorted element "bubbles up" to its correct
// position after each full pass.
//
// Algorithm walkthrough:
//
//	Input:  [5, 3, 8, 1]
//	Pass 1: [3, 5, 1, 8]  → 8 is now in place
//	Pass 2: [3, 1, 5, 8]  → 5 is now in place
//	Pass 3: [1, 3, 5, 8]  → sorted!
//
// Time Complexity:
//   - Best:    O(n)   — already sorted, early exit kicks in
//   - Average: O(n²)
//   - Worst:   O(n²)  — reverse sorted input
//
// Space Complexity: O(1) — in-place (on the copy)
//
// Stability: Stable — equal elements maintain relative order
func BubbleSort(arr []int) []int {
	// Work on a copy so we don't mutate the caller's slice.
	a := make([]int, len(arr))
	copy(a, arr)

	n := len(a)
	for i := 0; i < n-1; i++ {
		swapped := false

		// Each pass moves the largest remaining element to position n-i-1.
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j] // Go's clean multi-assign swap
				swapped = true
			}
		}

		// Optimization: if no swap happened, the array is already sorted.
		if !swapped {
			break
		}
	}

	return a
}
