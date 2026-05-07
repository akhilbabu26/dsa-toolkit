package sorting

// InsertionSort builds a sorted array one element at a time by taking each
// element from the unsorted portion and inserting it into its correct position
// in the already-sorted portion — similar to sorting a hand of playing cards.
//
// Algorithm walkthrough:
//
//	Input:  [5, 3, 8, 1]
//	i=1: key=3 → [3, 5, 8, 1]
//	i=2: key=8 → [3, 5, 8, 1]  (8 already in place)
//	i=3: key=1 → [1, 3, 5, 8]
//
// Time Complexity:
//   - Best:    O(n)   — nearly sorted input; inner loop rarely runs
//   - Average: O(n²)
//   - Worst:   O(n²)  — reverse sorted input
//
// Space Complexity: O(1)
//
// Stability: Stable
//
// Use when: input is small or nearly sorted — insertion sort outperforms
// quicksort in practice for n < ~10 elements.
func InsertionSort(arr []int) []int {
	a := make([]int, len(arr))
	copy(a, arr)

	for i := 1; i < len(a); i++ {
		// 'key' is the element we are currently trying to place.
		key := a[i]
		j := i - 1

		// Shift all elements greater than 'key' one position to the right
		// to make room for 'key'.
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}

		// Place 'key' in its correct sorted position.
		a[j+1] = key
	}

	return a
}
