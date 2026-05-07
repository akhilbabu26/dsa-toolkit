package bst

// This file implements the three classic DFS tree traversal strategies.
// Each traversal visits every node exactly once in a different order.
//
// Given this tree:
//
//	      4
//	     / \
//	    2   6
//	   / \ / \
//	  1  3 5  7
//
// Inorder   (L → Root → R): [1, 2, 3, 4, 5, 6, 7]  ← always sorted for a BST!
// Preorder  (Root → L → R): [4, 2, 1, 3, 6, 5, 7]  ← useful for tree cloning
// Postorder (L → R → Root): [1, 3, 2, 5, 7, 6, 4]  ← useful for tree deletion

// Inorder performs a Left → Node → Right depth-first traversal.
//
// Key insight: Inorder traversal of a BST always produces a sorted sequence.
// This is the most commonly used traversal for BSTs.
//
// Time Complexity:  O(n) — visits every node exactly once
// Space Complexity: O(h) — recursion stack, h = tree height
func (b *BST) Inorder() []int {
	var result []int
	inorder(b.Root, &result)
	return result
}

// inorder is the recursive helper. We pass a pointer to the result slice
// so each recursive call appends to the same underlying slice.
func inorder(node *TreeNode, result *[]int) {
	if node == nil {
		return // base case: empty subtree — nothing to visit
	}
	inorder(node.Left, result)   // 1. visit entire left subtree
	*result = append(*result, node.Value) // 2. visit this node
	inorder(node.Right, result)  // 3. visit entire right subtree
}

// Preorder performs a Node → Left → Right depth-first traversal.
//
// Key insight: The root is always the first element. Preorder is used to
// serialize/clone a tree — inserting values in preorder order recreates
// the same tree structure.
func (b *BST) Preorder() []int {
	var result []int
	preorder(b.Root, &result)
	return result
}

func preorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Value) // 1. visit this node first
	preorder(node.Left, result)           // 2. visit left subtree
	preorder(node.Right, result)          // 3. visit right subtree
}

// Postorder performs a Left → Right → Node depth-first traversal.
//
// Key insight: The root is always the LAST element. Postorder is used when
// you need to process children before parents — e.g., deleting a tree from
// memory, or evaluating expression trees bottom-up.
func (b *BST) Postorder() []int {
	var result []int
	postorder(b.Root, &result)
	return result
}

func postorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	postorder(node.Left, result)           // 1. visit left subtree
	postorder(node.Right, result)          // 2. visit right subtree
	*result = append(*result, node.Value)  // 3. visit this node last
}

// Height returns the height of the BST (number of edges on the longest path
// from root to a leaf). An empty tree has height -1; a single node has height 0.
//
// Time Complexity: O(n) — must visit every node
func (b *BST) Height() int {
	return height(b.Root)
}

func height(node *TreeNode) int {
	if node == nil {
		return -1 // empty subtree contributes -1 to edge count
	}
	leftH := height(node.Left)
	rightH := height(node.Right)
	// Height = 1 edge to tallest child + child's height
	if leftH > rightH {
		return 1 + leftH
	}
	return 1 + rightH
}
