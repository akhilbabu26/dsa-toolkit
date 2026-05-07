package sorting

// SelectionSort divides the array into two parts: sorted (left) and unsorted
// (right). In each pass, it finds the minimum element in the unsorted part
// and swaps it with the first unsorted element.
//
// Algorithm walkthrough:
//
//	Input:  [5, 3, 8, 1]
//	Pass 1: min=1 at idx 3 → swap with idx 0 → [1, 3, 8, 5]
//	Pass 2: min=3 at idx 1 → already in place  → [1, 3, 8, 5]
//	Pass 3: min=5 at idx 3 → swap with idx 2 → [1, 3, 5, 8]
//
// Time Complexity:
//   - Best:    O(n²)  — always scans the unsorted part regardless
//   - Average: O(n²)
//   - Worst:   O(n²)
//
// Space Complexity: O(1)
//
// Stability: NOT stable — a swap can change the relative order of equal elements
//
// Trade-off: Selection sort makes at most O(n) swaps — useful when memory
// writes are expensive (e.g., flash memory).
func SelectionSort(arr []int) []int {
	a := make([]int, len(arr))
	copy(a, arr)

	n := len(a)
	for i := 0; i < n-1; i++ {
		// Assume the current position holds the minimum.
		minIdx := i

		// Scan the rest of the array to find the true minimum.
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}

		// Swap the found minimum with the first unsorted element.
		a[i], a[minIdx] = a[minIdx], a[i]
	}

	return a
}
