package heap

import "testing"

// ── MinHeap Tests ─────────────────────────────────────────────────────────────

func TestMinHeapInsertAndPeek(t *testing.T) {
	h := NewMinHeap()

	if !h.IsEmpty() {
		t.Error("new heap should be empty")
	}

	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)
	h.Insert(9)

	min, err := h.Peek()
	if err != nil {
		t.Fatalf("Peek error: %v", err)
	}
	if min != 1 {
		t.Errorf("MinHeap Peek() = %d; want 1", min)
	}
	// Peek should not remove the element
	if h.Size() != 5 {
		t.Errorf("size after Peek() = %d; want 5", h.Size())
	}
}

func TestMinHeapExtractMin(t *testing.T) {
	h := NewMinHeap()
	values := []int{5, 3, 8, 1, 9, 2, 7}
	for _, v := range values {
		h.Insert(v)
	}

	// Extracting all elements should yield them in ascending order.
	expected := []int{1, 2, 3, 5, 7, 8, 9}
	for _, want := range expected {
		got, err := h.ExtractMin()
		if err != nil {
			t.Fatalf("ExtractMin error: %v", err)
		}
		if got != want {
			t.Errorf("ExtractMin() = %d; want %d", got, want)
		}
	}

	if !h.IsEmpty() {
		t.Error("heap should be empty after extracting all elements")
	}
}

func TestMinHeapSingleElement(t *testing.T) {
	h := NewMinHeap()
	h.Insert(42)

	min, err := h.ExtractMin()
	if err != nil || min != 42 {
		t.Errorf("single element ExtractMin() = %d, %v; want 42, nil", min, err)
	}
	if !h.IsEmpty() {
		t.Error("heap should be empty")
	}
}

func TestMinHeapEmptyPeek(t *testing.T) {
	h := NewMinHeap()
	_, err := h.Peek()
	if err == nil {
		t.Error("Peek on empty heap should return error")
	}
}

func TestMinHeapEmptyExtract(t *testing.T) {
	h := NewMinHeap()
	_, err := h.ExtractMin()
	if err == nil {
		t.Error("ExtractMin on empty heap should return error")
	}
}

func TestMinHeapDuplicates(t *testing.T) {
	h := NewMinHeap()
	for _, v := range []int{3, 3, 1, 1, 2} {
		h.Insert(v)
	}
	got, _ := h.ExtractMin()
	if got != 1 {
		t.Errorf("ExtractMin with duplicates = %d; want 1", got)
	}
}

// ── MaxHeap Tests ─────────────────────────────────────────────────────────────

func TestMaxHeapInsertAndPeek(t *testing.T) {
	h := NewMaxHeap()
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)
	h.Insert(9)

	max, err := h.Peek()
	if err != nil {
		t.Fatalf("Peek error: %v", err)
	}
	if max != 9 {
		t.Errorf("MaxHeap Peek() = %d; want 9", max)
	}
}

func TestMaxHeapExtractMax(t *testing.T) {
	h := NewMaxHeap()
	values := []int{5, 3, 8, 1, 9, 2, 7}
	for _, v := range values {
		h.Insert(v)
	}

	// Extracting all elements should yield them in descending order.
	expected := []int{9, 8, 7, 5, 3, 2, 1}
	for _, want := range expected {
		got, err := h.ExtractMax()
		if err != nil {
			t.Fatalf("ExtractMax error: %v", err)
		}
		if got != want {
			t.Errorf("ExtractMax() = %d; want %d", got, want)
		}
	}
}

func TestMaxHeapEmptyErrors(t *testing.T) {
	h := NewMaxHeap()
	if _, err := h.Peek(); err == nil {
		t.Error("Peek on empty MaxHeap should error")
	}
	if _, err := h.ExtractMax(); err == nil {
		t.Error("ExtractMax on empty MaxHeap should error")
	}
}
