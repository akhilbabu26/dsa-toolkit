package bst

import (
	"reflect"
	"testing"
)

// buildTree is a test helper that inserts values and returns the BST.
func buildTree(values []int) *BST {
	tree := New()
	for _, v := range values {
		tree.Insert(v)
	}
	return tree
}

// ── Insert & Search Tests ─────────────────────────────────────────────────────

func TestInsertAndSearch(t *testing.T) {
	tree := buildTree([]int{5, 3, 7, 1, 4, 6, 8})

	presentValues := []int{5, 3, 7, 1, 4, 6, 8}
	for _, v := range presentValues {
		if !tree.Search(v) {
			t.Errorf("Search(%d) = false; want true", v)
		}
	}

	absentValues := []int{0, 2, 9, 100}
	for _, v := range absentValues {
		if tree.Search(v) {
			t.Errorf("Search(%d) = true; want false", v)
		}
	}
}

func TestInsertDuplicateIgnored(t *testing.T) {
	tree := buildTree([]int{5, 5, 5})
	result := tree.Inorder()
	if !reflect.DeepEqual(result, []int{5}) {
		t.Errorf("duplicate inserts should yield single node; got %v", result)
	}
}

func TestSearchEmptyTree(t *testing.T) {
	tree := New()
	if tree.Search(42) {
		t.Error("Search on empty tree should return false")
	}
}

// ── Traversal Tests ───────────────────────────────────────────────────────────

func TestInorder(t *testing.T) {
	//      5
	//     / \
	//    3   7
	//   / \ / \
	//  1  4 6  8
	tree := buildTree([]int{5, 3, 7, 1, 4, 6, 8})
	got := tree.Inorder()
	want := []int{1, 3, 4, 5, 6, 7, 8} // must be sorted
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Inorder() = %v; want %v", got, want)
	}
}

func TestPreorder(t *testing.T) {
	tree := buildTree([]int{5, 3, 7, 1, 4, 6, 8})
	got := tree.Preorder()
	want := []int{5, 3, 1, 4, 7, 6, 8} // root first
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Preorder() = %v; want %v", got, want)
	}
}

func TestPostorder(t *testing.T) {
	tree := buildTree([]int{5, 3, 7, 1, 4, 6, 8})
	got := tree.Postorder()
	want := []int{1, 4, 3, 6, 8, 7, 5} // root last
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Postorder() = %v; want %v", got, want)
	}
}

func TestTraversalEmptyTree(t *testing.T) {
	tree := New()
	if tree.Inorder() != nil {
		t.Error("Inorder of empty tree should return nil slice")
	}
}

// ── Delete Tests ──────────────────────────────────────────────────────────────

func TestDeleteLeaf(t *testing.T) {
	tree := buildTree([]int{5, 3, 7})
	tree.Delete(3) // 3 is a leaf
	if tree.Search(3) {
		t.Error("leaf node 3 should be deleted")
	}
	got := tree.Inorder()
	if !reflect.DeepEqual(got, []int{5, 7}) {
		t.Errorf("after deleting leaf 3: got %v, want [5 7]", got)
	}
}

func TestDeleteOneChild(t *testing.T) {
	tree := buildTree([]int{5, 3, 7, 2})
	tree.Delete(3) // 3 has one child (2)
	if tree.Search(3) {
		t.Error("node 3 should be deleted")
	}
	if !tree.Search(2) {
		t.Error("child 2 should still exist after deleting parent 3")
	}
}

func TestDeleteTwoChildren(t *testing.T) {
	tree := buildTree([]int{5, 3, 7, 1, 4, 6, 8})
	tree.Delete(3) // 3 has two children (1 and 4)
	if tree.Search(3) {
		t.Error("node 3 should be deleted")
	}
	// Inorder must still be sorted after two-child deletion
	got := tree.Inorder()
	want := []int{1, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("after deleting 3: got %v, want %v", got, want)
	}
}

func TestDeleteRoot(t *testing.T) {
	tree := buildTree([]int{5, 3, 7})
	tree.Delete(5) // delete root
	if tree.Search(5) {
		t.Error("root 5 should be deleted")
	}
	got := tree.Inorder()
	want := []int{3, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("after deleting root: got %v, want %v", got, want)
	}
}

func TestDeleteNonExistent(t *testing.T) {
	tree := buildTree([]int{5, 3, 7})
	tree.Delete(99) // should be a no-op
	got := tree.Inorder()
	want := []int{3, 5, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("deleting non-existent should be no-op: got %v, want %v", got, want)
	}
}

// ── Min / Max / Height Tests ──────────────────────────────────────────────────

func TestMinMax(t *testing.T) {
	tree := buildTree([]int{5, 3, 7, 1, 9})

	min, ok := tree.Min()
	if !ok || min != 1 {
		t.Errorf("Min() = %d, %v; want 1, true", min, ok)
	}

	max, ok := tree.Max()
	if !ok || max != 9 {
		t.Errorf("Max() = %d, %v; want 9, true", max, ok)
	}
}

func TestMinMaxEmptyTree(t *testing.T) {
	tree := New()
	if _, ok := tree.Min(); ok {
		t.Error("Min of empty tree should return ok=false")
	}
	if _, ok := tree.Max(); ok {
		t.Error("Max of empty tree should return ok=false")
	}
}

func TestHeight(t *testing.T) {
	tree := New()
	if tree.Height() != -1 {
		t.Errorf("empty tree height: got %d, want -1", tree.Height())
	}

	tree.Insert(5)
	if tree.Height() != 0 {
		t.Errorf("single node height: got %d, want 0", tree.Height())
	}

	tree.Insert(3)
	tree.Insert(7)
	if tree.Height() != 1 {
		t.Errorf("height of 3-node balanced tree: got %d, want 1", tree.Height())
	}
}
