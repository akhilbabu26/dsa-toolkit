// Package heap implements MinHeap and MaxHeap using slice-backed arrays.
//
// # How a Heap works (beginner primer)
//
// A heap is a complete binary tree stored compactly in a slice.
// For a node at index i:
//   - Parent:       (i-1) / 2
//   - Left child:   2*i + 1
//   - Right child:  2*i + 2
//
// MinHeap property: every parent ≤ its children → root is always the minimum.
//
//	Heap slice: [1, 3, 5, 7, 9, 8]
//
//	Visualised:
//	        1          ← root (minimum)
//	       / \
//	      3   5
//	     / \ /
//	    7  9 8
//
// Operations:
//   - Insert: append to end, then "heapify up" (swap with parent if smaller)
//   - ExtractMin: swap root with last element, shrink slice, "heapify down"
//   - Peek: return root element without removing it — O(1)
package heap

import "fmt"

// MinHeap is a min-heap data structure backed by a slice.
// The element at index 0 is always the smallest value.
type MinHeap struct {
	data []int
}

// NewMinHeap creates and returns an empty MinHeap.
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// Size returns the number of elements in the heap.
func (h *MinHeap) Size() int { return len(h.data) }

// IsEmpty reports whether the heap has no elements.
func (h *MinHeap) IsEmpty() bool { return len(h.data) == 0 }

// Peek returns the minimum element without removing it.
// Returns an error if the heap is empty.
//
// Time Complexity: O(1) — root is always at index 0
func (h *MinHeap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, fmt.Errorf("minheap: peek on empty heap")
	}
	return h.data[0], nil
}

// Insert adds a value into the heap and restores the heap property.
//
// Steps:
//  1. Append the new element to the end of the slice.
//  2. "Heapify up": compare with parent; swap if smaller; repeat.
//
// Time Complexity: O(log n) — at most h swaps where h = height = log₂(n)
func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

// heapifyUp bubbles the element at index i toward the root until the heap
// property is restored (parent ≤ child at every level).
func (h *MinHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2 // integer division gives parent index
		if h.data[i] < h.data[parent] {
			// Child is smaller than parent — swap to restore min-heap property.
			h.data[i], h.data[parent] = h.data[parent], h.data[i]
			i = parent // move up and check again
		} else {
			break // heap property satisfied
		}
	}
}

// ExtractMin removes and returns the minimum element (the root).
//
// Steps:
//  1. Save the root value (the minimum).
//  2. Move the last element to the root position.
//  3. Shrink the slice by one.
//  4. "Heapify down": push the new root down until heap property is restored.
//
// Time Complexity: O(log n)
func (h *MinHeap) ExtractMin() (int, error) {
	if h.IsEmpty() {
		return 0, fmt.Errorf("minheap: extract from empty heap")
	}

	min := h.data[0]              // save root
	last := len(h.data) - 1
	h.data[0] = h.data[last]     // move last element to root
	h.data = h.data[:last]        // shrink slice (discard last position)
	h.heapifyDown(0)              // restore heap property from root down

	return min, nil
}

// heapifyDown pushes the element at index i downward until the heap property
// is restored. At each step it swaps with the smaller of the two children.
func (h *MinHeap) heapifyDown(i int) {
	n := len(h.data)
	for {
		smallest := i
		left := 2*i + 1  // left child index
		right := 2*i + 2 // right child index

		// Check if left child exists and is smaller than current smallest
		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}
		// Check if right child exists and is smaller than current smallest
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == i {
			break // both children are larger (or don't exist) — we're done
		}

		// Swap with the smallest child and continue downward.
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

// Snapshot returns a copy of the internal slice for inspection/testing.
// Do NOT modify the returned slice directly.
func (h *MinHeap) Snapshot() []int {
	cp := make([]int, len(h.data))
	copy(cp, h.data)
	return cp
}
