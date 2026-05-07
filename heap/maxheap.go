package heap

import "fmt"

// MaxHeap is a max-heap data structure backed by a slice.
// The element at index 0 is always the largest value.
//
// MaxHeap is the mirror image of MinHeap — the only difference is that
// comparisons are reversed: parent ≥ children (instead of parent ≤ children).
//
// Use cases:
//   - Priority queues where high-priority items are processed first
//   - HeapSort (sort in ascending order using a max-heap)
//   - Finding the k largest elements efficiently
type MaxHeap struct {
	data []int
}

// NewMaxHeap creates and returns an empty MaxHeap.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Size returns the number of elements in the heap.
func (h *MaxHeap) Size() int { return len(h.data) }

// IsEmpty reports whether the heap has no elements.
func (h *MaxHeap) IsEmpty() bool { return len(h.data) == 0 }

// Peek returns the maximum element without removing it.
// Returns an error if the heap is empty.
//
// Time Complexity: O(1)
func (h *MaxHeap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, fmt.Errorf("maxheap: peek on empty heap")
	}
	return h.data[0], nil
}

// Insert adds a value into the max-heap and restores the heap property.
//
// Time Complexity: O(log n)
func (h *MaxHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

// heapifyUp bubbles the element at index i up toward the root until the
// max-heap property (parent ≥ children) is satisfied.
func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[i] > h.data[parent] {
			// Child is GREATER than parent — swap (MaxHeap: parent must be ≥ child)
			h.data[i], h.data[parent] = h.data[parent], h.data[i]
			i = parent
		} else {
			break
		}
	}
}

// ExtractMax removes and returns the maximum element (the root).
//
// Time Complexity: O(log n)
func (h *MaxHeap) ExtractMax() (int, error) {
	if h.IsEmpty() {
		return 0, fmt.Errorf("maxheap: extract from empty heap")
	}

	max := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	h.heapifyDown(0)

	return max, nil
}

// heapifyDown pushes the element at index i downward, always swapping with
// the LARGER of the two children (opposite of MinHeap which uses the smaller).
func (h *MaxHeap) heapifyDown(i int) {
	n := len(h.data)
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < n && h.data[right] > h.data[largest] {
			largest = right
		}

		if largest == i {
			break // heap property satisfied
		}

		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}

// Snapshot returns a copy of the internal slice for inspection/testing.
func (h *MaxHeap) Snapshot() []int {
	cp := make([]int, len(h.data))
	copy(cp, h.data)
	return cp
}
