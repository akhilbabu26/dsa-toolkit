// Package bst implements a Binary Search Tree (BST) with insert, search,
// delete, and three traversal strategies.
//
// # What is a BST?
//
// A Binary Search Tree is a rooted binary tree where for every node N:
//   - All values in N's LEFT subtree are LESS than N.Value
//   - All values in N's RIGHT subtree are GREATER than N.Value
//
// This ordering property makes search, insert, and delete all O(h) where h
// is the tree height — O(log n) for balanced trees, O(n) for degenerate ones.
//
//	        8
//	       / \
//	      3   10
//	     / \    \
//	    1   6    14
//	       / \   /
//	      4   7 13
package bst

// TreeNode represents a single node in the BST.
type TreeNode struct {
	Value int       // the data stored in this node
	Left  *TreeNode // pointer to the left child (smaller values)
	Right *TreeNode // pointer to the right child (larger values)
}

// BST is the top-level Binary Search Tree structure.
// It holds a pointer to the root node of the tree.
type BST struct {
	Root *TreeNode
}

// New creates and returns an empty Binary Search Tree.
func New() *BST {
	return &BST{}
}

// Insert adds a new value into the BST while maintaining BST ordering.
//
// Algorithm:
//  1. Start at the root.
//  2. If value < current node → go left.
//  3. If value > current node → go right.
//  4. If we reach nil, place the new node here.
//  5. Duplicates are ignored (set-like behaviour).
//
// Time Complexity: O(h) — O(log n) average, O(n) worst
// Space Complexity: O(h) — call stack depth due to recursion
func (b *BST) Insert(value int) {
	b.Root = insertNode(b.Root, value)
}

// insertNode is the recursive helper for Insert. It returns the (possibly new)
// root of the subtree rooted at 'node'.
func insertNode(node *TreeNode, value int) *TreeNode {
	// Base case: we've reached an empty spot — create and return the new node.
	if node == nil {
		return &TreeNode{Value: value}
	}

	switch {
	case value < node.Value:
		// Value belongs in the left subtree — recurse left.
		node.Left = insertNode(node.Left, value)
	case value > node.Value:
		// Value belongs in the right subtree — recurse right.
		node.Right = insertNode(node.Right, value)
	// value == node.Value: duplicate — do nothing (BST stores unique values)
	}

	return node
}

// Search checks whether a value exists in the BST.
//
// Time Complexity: O(h)
func (b *BST) Search(value int) bool {
	return searchNode(b.Root, value)
}

func searchNode(node *TreeNode, value int) bool {
	if node == nil {
		return false // value is not in the tree
	}
	switch {
	case value == node.Value:
		return true
	case value < node.Value:
		return searchNode(node.Left, value)
	default:
		return searchNode(node.Right, value)
	}
}

// Delete removes a value from the BST, handling all three structural cases:
//
//  1. Leaf node (no children): simply remove it.
//  2. One child: replace node with its only child.
//  3. Two children: replace node's value with its in-order successor
//     (the smallest value in the RIGHT subtree), then delete the successor.
//
// Time Complexity: O(h)
func (b *BST) Delete(value int) {
	b.Root = deleteNode(b.Root, value)
}

func deleteNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil // value not found — nothing to do
	}

	if value < node.Value {
		// Target is in the left subtree.
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		// Target is in the right subtree.
		node.Right = deleteNode(node.Right, value)
	} else {
		// ── Found the node to delete ──────────────────────────────────────

		// Case 1 & 2: no left child — promote right child (may be nil).
		if node.Left == nil {
			return node.Right
		}
		// Case 2: no right child — promote left child.
		if node.Right == nil {
			return node.Left
		}

		// Case 3: two children.
		// Find the in-order successor: smallest node in the right subtree.
		// It is guaranteed to have at most one child (no left child).
		successor := minNode(node.Right)

		// Replace this node's value with the successor's value.
		node.Value = successor.Value

		// Delete the successor from the right subtree (it's now duplicated).
		node.Right = deleteNode(node.Right, successor.Value)
	}

	return node
}

// minNode returns the leftmost (minimum) node in a subtree.
// Used by deleteNode to find the in-order successor.
func minNode(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Min returns the minimum value stored in the BST.
// Returns 0 and false if the tree is empty.
func (b *BST) Min() (int, bool) {
	if b.Root == nil {
		return 0, false
	}
	return minNode(b.Root).Value, true
}

// Max returns the maximum value stored in the BST.
// Returns 0 and false if the tree is empty.
func (b *BST) Max() (int, bool) {
	if b.Root == nil {
		return 0, false
	}
	node := b.Root
	for node.Right != nil {
		node = node.Right
	}
	return node.Value, true
}
