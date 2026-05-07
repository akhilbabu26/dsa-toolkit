<div align="center">

# 🚀 Go DSA Toolkit

**A modular, beginner-friendly Data Structures & Algorithms library in Go**

![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![Tests](https://img.shields.io/badge/Tests-Passing-brightgreen?style=flat-square)

</div>

---

## 📖 Project Overview

**Go DSA Toolkit** is a clean, well-commented implementation of core Data Structures & Algorithms in Go. It is designed to be:

- 🧠 **Educational** — every function explains *why*, not just *what*
- 🔰 **Beginner-friendly** — step-by-step algorithm walkthroughs in comments
- 🧩 **Modular** — each topic is an independent, importable Go package
- ✅ **Tested** — table-driven unit tests with edge cases for every algorithm
- 🖥️ **Interactive** — a CLI runner demonstrates all algorithms in your terminal

---

## 📁 Folder Structure

```
dsa-toolkit/
│
├── main.go              ← Interactive CLI runner (go run main.go)
├── go.mod               ← Go module definition
├── README.md            ← This file
│
├── sorting/             ← Sorting algorithms
│   ├── bubble.go        ← Bubble Sort
│   ├── insertion.go     ← Insertion Sort
│   ├── selection.go     ← Selection Sort
│   └── sorting_test.go  ← Table-driven tests (8 cases each)
│
├── stringsutil/         ← UTF-8-safe string utilities
│   ├── reverse.go       ← Unicode-correct string reversal
│   ├── palindrome.go    ← Palindrome checker (case + punctuation aware)
│   ├── vowels.go        ← Vowel counter with frequency map
│   └── strings_test.go  ← Tests including emoji and accented chars
│
├── hashmap/             ← Map-based data structures
│   ├── frequency.go     ← Word frequency counter & duplicate detector
│   ├── set.go           ← Generic Set with Union & Intersection
│   └── hashmap_test.go  ← Tests for all map operations
│
├── bst/                 ← Binary Search Tree
│   ├── bst.go           ← Insert, Search, Delete (all 3 cases), Min, Max
│   ├── traversal.go     ← Inorder, Preorder, Postorder, Height
│   └── bst_test.go      ← 15+ test cases including edge cases
│
├── heap/                ← Heap data structures
│   ├── minheap.go       ← Min-Heap with full heapify logic
│   ├── maxheap.go       ← Max-Heap (mirror of MinHeap)
│   └── heap_test.go     ← Extract ordering, empty errors, duplicates
│
└── examples/
    └── sample_output.txt ← Pre-generated CLI output for all sections
```

---

## ⚙️ Implemented Algorithms & Data Structures

### 🔢 Sorting

| Algorithm      | Strategy                        | Stable |
|----------------|---------------------------------|--------|
| Bubble Sort    | Swap adjacent out-of-order pairs | ✅     |
| Insertion Sort | Insert each element into place  | ✅     |
| Selection Sort | Find min, place at front        | ❌     |

### 🔤 String Utilities

| Function        | Description                                      |
|-----------------|--------------------------------------------------|
| `Reverse`       | UTF-8 safe reversal using `[]rune` conversion    |
| `IsPalindrome`  | Ignores case, spaces, punctuation                |
| `CountVowels`   | Returns count + per-vowel frequency map          |

### 🗂️ HashMap

| Function/Type      | Description                                   |
|--------------------|-----------------------------------------------|
| `WordFrequency`    | Count occurrences of each word (normalised)   |
| `CharFrequency`    | Count rune occurrences in a string            |
| `FindDuplicates`   | Return words appearing more than once         |
| `Set`              | Add, Remove, Contains, Union, Intersection    |

### 🌳 Binary Search Tree

| Operation        | Description                                           |
|------------------|-------------------------------------------------------|
| `Insert`         | Recursive insert maintaining BST ordering             |
| `Search`         | O(h) search using BST property                        |
| `Delete`         | Handles leaf / one-child / two-children cases         |
| `Inorder`        | Left → Node → Right (always sorted for BST)           |
| `Preorder`       | Node → Left → Right (root first)                      |
| `Postorder`      | Left → Right → Node (root last)                       |
| `Min` / `Max`    | Leftmost / rightmost node                             |
| `Height`         | Length of longest root-to-leaf path                   |

### 🏔️ Heap

| Operation      | MinHeap             | MaxHeap             |
|----------------|---------------------|---------------------|
| `Insert`       | O(log n)            | O(log n)            |
| `ExtractMin/Max` | O(log n)          | O(log n)            |
| `Peek`         | O(1)                | O(1)                |
| `HeapifyUp`    | Bubble up after insert | Bubble up after insert |
| `HeapifyDown`  | Sink down after extract | Sink down after extract |

---

## 🚀 How to Run

### Prerequisites

- [Go 1.22+](https://go.dev/dl/) installed
- `git clone` the repository

### Clone & Run

```bash
# Clone the repository
git clone https://github.com/akhilbabu26/dsa-toolkit.git
cd dsa-toolkit

# Run the interactive CLI
go run main.go
```

You'll see an interactive menu:

```
╔══════════════════════════════════════╗
║      🚀  Go DSA Playground  🚀       ║
╠══════════════════════════════════════╣
║  1. Sorting Algorithms               ║
║  2. String Utilities                 ║
║  3. HashMap / Set                    ║
║  4. Binary Search Tree               ║
║  5. Heap (Min & Max)                 ║
║  6. Exit                             ║
╚══════════════════════════════════════╝
Enter your choice:
```

---

## 🧪 How to Test

### Run all tests at once

```bash
go test ./...
```

### Run tests for a specific module

```bash
go test ./sorting/...     # Sorting tests
go test ./stringsutil/... # String utility tests
go test ./hashmap/...     # HashMap & Set tests
go test ./bst/...         # BST tests
go test ./heap/...        # Heap tests
```

### Run tests with verbose output

```bash
go test ./... -v
```

### Run tests with coverage

```bash
go test ./... -cover
```

---

## 💡 Example Outputs

### Sorting

```
Input:          [64 34 25 12 22 11 90]
Bubble Sort:    [11 12 22 25 34 64 90]
Insertion Sort: [11 12 22 25 34 64 90]
Selection Sort: [11 12 22 25 34 64 90]
```

### BST Traversals

```
Inserting: [8 3 10 1 6 14 4 7 13]

Inorder   (sorted):     [1 3 4 6 7 8 10 13 14]
Preorder  (root first): [8 3 1 6 4 7 10 14 13]
Postorder (root last):  [1 4 7 6 3 13 14 10 8]

Min: 1 | Max: 14 | Height: 3
```

### Heap

```
Input: [5 3 8 1 9 2 7 4 6]

MinHeap — Extract order: [1, 2, 3, 4, 5, 6, 7, 8, 9]  ← ascending
MaxHeap — Extract order: [9, 8, 7, 6, 5, 4, 3, 2, 1]  ← descending
```

> See [`examples/sample_output.txt`](examples/sample_output.txt) for the full output of all sections.

---

## ⏱️ Time Complexity Reference

### Sorting

| Algorithm      | Best   | Average | Worst  | Space |
|----------------|--------|---------|--------|-------|
| Bubble Sort    | O(n)   | O(n²)   | O(n²)  | O(1)  |
| Insertion Sort | O(n)   | O(n²)   | O(n²)  | O(1)  |
| Selection Sort | O(n²)  | O(n²)   | O(n²)  | O(1)  |

### Data Structures

| Structure  | Insert    | Search    | Delete    | Peek |
|------------|-----------|-----------|-----------|------|
| BST (avg)  | O(log n)  | O(log n)  | O(log n)  | —    |
| BST (worst)| O(n)      | O(n)      | O(n)      | —    |
| MinHeap    | O(log n)  | O(n)      | O(log n)  | O(1) |
| MaxHeap    | O(log n)  | O(n)      | O(log n)  | O(1) |
| HashMap    | O(1) avg  | O(1) avg  | O(1) avg  | —    |
| Set        | O(1) avg  | O(1) avg  | O(1) avg  | —    |

---

## 🌱 Learning Goals

After exploring this project you will understand:

- ✅ How to organise a Go project using packages and modules
- ✅ Why `[]rune` is necessary for safe Unicode string manipulation
- ✅ How recursive BST operations work (insert, delete, traversals)
- ✅ Why `map[T]struct{}` is more memory-efficient than `map[T]bool`
- ✅ How heaps use a slice to represent a complete binary tree
- ✅ How table-driven tests keep test code DRY and maintainable
- ✅ The trade-offs between different sorting algorithms
- ✅ The importance of edge case handling (empty input, duplicates, nil)

---

## 🔮 Future Improvements

- [ ] **Linked List** — Singly, Doubly, Circular
- [ ] **Graph Algorithms** — BFS, DFS, Dijkstra's shortest path
- [ ] **Dynamic Programming** — Fibonacci, Knapsack, LCS
- [ ] **Trie** — Prefix tree for autocomplete
- [ ] **AVL / Red-Black Tree** — Self-balancing BST
- [ ] **Merge Sort & Quick Sort** — O(n log n) sorting
- [ ] **Benchmark tests** — `go test -bench` profiling
- [ ] **GitHub Actions CI** — Automatic test runs on push

---

## 📜 License

This project is open source and available under the [MIT License](LICENSE).

---

<div align="center">
Made with ❤️ and Go &nbsp;|&nbsp; Happy Learning! 🎓
</div>
