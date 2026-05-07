// Go DSA Toolkit — main.go
//
// This is the CLI entry point for the DSA Toolkit.
// It provides an interactive menu that demonstrates every algorithm and
// data structure implemented in this repository.
//
// Run with:
//
//	go run main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/akhilbabu26/dsa-toolkit/bst"
	"github.com/akhilbabu26/dsa-toolkit/hashmap"
	"github.com/akhilbabu26/dsa-toolkit/heap"
	sorting "github.com/akhilbabu26/dsa-toolkit/sorting"
	"github.com/akhilbabu26/dsa-toolkit/stringsutil"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printBanner()
		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			demoSorting()
		case "2":
			demoStrings()
		case "3":
			demoHashMap()
		case "4":
			demoBST()
		case "5":
			demoHeap()
		case "6":
			fmt.Println("\n👋  Thanks for using Go DSA Toolkit. Happy coding!")
			os.Exit(0)
		default:
			fmt.Println("\n⚠️  Invalid choice. Please enter a number between 1 and 6.")
		}

		fmt.Println("\nPress ENTER to return to the menu...")
		reader.ReadString('\n')
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Banner
// ─────────────────────────────────────────────────────────────────────────────

func printBanner() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║      🚀  Go DSA Playground  🚀       ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  1. Sorting Algorithms               ║")
	fmt.Println("║  2. String Utilities                 ║")
	fmt.Println("║  3. HashMap / Set                    ║")
	fmt.Println("║  4. Binary Search Tree               ║")
	fmt.Println("║  5. Heap (Min & Max)                 ║")
	fmt.Println("║  6. Exit                             ║")
	fmt.Println("╚══════════════════════════════════════╝")
}

// ─────────────────────────────────────────────────────────────────────────────
// 1. Sorting Demo
// ─────────────────────────────────────────────────────────────────────────────

func demoSorting() {
	section("Sorting Algorithms")
	input := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("  Input:          %v\n\n", input)
	fmt.Printf("  Bubble Sort:    %v\n", sorting.BubbleSort(input))
	fmt.Printf("  Insertion Sort: %v\n", sorting.InsertionSort(input))
	fmt.Printf("  Selection Sort: %v\n", sorting.SelectionSort(input))

	fmt.Println()
	fmt.Println("  Complexity Summary:")
	fmt.Println("  ┌─────────────────┬──────────┬──────────┬──────────┬─────────┐")
	fmt.Println("  │ Algorithm       │ Best     │ Average  │ Worst    │ Space   │")
	fmt.Println("  ├─────────────────┼──────────┼──────────┼──────────┼─────────┤")
	fmt.Println("  │ Bubble Sort     │ O(n)     │ O(n²)    │ O(n²)    │ O(1)    │")
	fmt.Println("  │ Insertion Sort  │ O(n)     │ O(n²)    │ O(n²)    │ O(1)    │")
	fmt.Println("  │ Selection Sort  │ O(n²)    │ O(n²)    │ O(n²)    │ O(1)    │")
	fmt.Println("  └─────────────────┴──────────┴──────────┴──────────┴─────────┘")
}

// ─────────────────────────────────────────────────────────────────────────────
// 2. String Utilities Demo
// ─────────────────────────────────────────────────────────────────────────────

func demoStrings() {
	section("String Utilities")

	words := []string{"hello", "racecar", "café", "A man, a plan, a canal: Panama", "🙂😎"}
	for _, w := range words {
		rev := stringsutil.Reverse(w)
		pal := stringsutil.IsPalindrome(w)
		fmt.Printf("  Input:     %q\n", w)
		fmt.Printf("  Reversed:  %q\n", rev)
		fmt.Printf("  Palindrome: %v  (cleaned: %q)\n\n", pal.IsPalin, pal.Cleaned)
	}

	vowelInputs := []string{"Hello, World!", "rhythm", "Beautiful"}
	fmt.Println("  Vowel Counter:")
	for _, s := range vowelInputs {
		r := stringsutil.CountVowels(s)
		// Sort frequency output for consistent display
		type kv struct {
			k rune
			v int
		}
		var pairs []kv
		for k, v := range r.Frequency {
			pairs = append(pairs, kv{k, v})
		}
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].k < pairs[j].k })
		fmt.Printf("  %q → vowels: %d, freq: %v\n", s, r.Count, pairs)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// 3. HashMap Demo
// ─────────────────────────────────────────────────────────────────────────────

func demoHashMap() {
	section("HashMap & Set")

	// Word frequency
	words := strings.Fields("the quick brown fox jumps over the lazy dog the fox")
	freq := hashmap.WordFrequency(words)
	fmt.Println("  Word Frequency (\"the quick brown fox ... the fox\"):")
	// Collect and sort keys for deterministic output
	keys := make([]string, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		bar := strings.Repeat("█", freq[k])
		fmt.Printf("    %-10s %s (%d)\n", k, bar, freq[k])
	}

	// Duplicates
	fmt.Println()
	dups := hashmap.FindDuplicates(words)
	sort.Strings(dups)
	fmt.Printf("  Duplicate words: %v\n", dups)

	// Set operations
	fmt.Println()
	fmt.Println("  Set Operations:")
	a := hashmap.NewSet()
	for _, lang := range []string{"Go", "Python", "Rust"} {
		a.Add(lang)
	}
	b := hashmap.NewSet()
	for _, lang := range []string{"Python", "Java", "Go"} {
		b.Add(lang)
	}

	u := hashmap.Union(a, b)
	i := hashmap.Intersection(a, b)

	uElems := u.Elements()
	iElems := i.Elements()
	sort.Strings(uElems)
	sort.Strings(iElems)

	fmt.Printf("  Set A:            %v\n", sortedElems(a))
	fmt.Printf("  Set B:            %v\n", sortedElems(b))
	fmt.Printf("  Union A ∪ B:      %v\n", uElems)
	fmt.Printf("  Intersection A ∩ B: %v\n", iElems)
}

func sortedElems(s *hashmap.Set) []string {
	e := s.Elements()
	sort.Strings(e)
	return e
}

// ─────────────────────────────────────────────────────────────────────────────
// 4. BST Demo
// ─────────────────────────────────────────────────────────────────────────────

func demoBST() {
	section("Binary Search Tree")

	tree := bst.New()
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	fmt.Printf("  Inserting: %v\n\n", values)
	for _, v := range values {
		tree.Insert(v)
	}

	fmt.Printf("  Inorder   (sorted):         %v\n", tree.Inorder())
	fmt.Printf("  Preorder  (root first):     %v\n", tree.Preorder())
	fmt.Printf("  Postorder (root last):      %v\n", tree.Postorder())

	min, _ := tree.Min()
	max, _ := tree.Max()
	fmt.Printf("\n  Min: %d | Max: %d | Height: %d\n", min, max, tree.Height())

	fmt.Println()
	fmt.Printf("  Search(6):  %v\n", tree.Search(6))
	fmt.Printf("  Search(99): %v\n", tree.Search(99))

	// Delete demo
	fmt.Println()
	fmt.Println("  Delete 6 (two children), then 10 (one child), then 1 (leaf):")
	tree.Delete(6)
	tree.Delete(10)
	tree.Delete(1)
	fmt.Printf("  Inorder after deletions: %v\n", tree.Inorder())
}

// ─────────────────────────────────────────────────────────────────────────────
// 5. Heap Demo
// ─────────────────────────────────────────────────────────────────────────────

func demoHeap() {
	section("Heap (Min & Max)")

	values := []int{5, 3, 8, 1, 9, 2, 7, 4, 6}
	fmt.Printf("  Input: %v\n\n", values)

	// MinHeap
	minH := heap.NewMinHeap()
	for _, v := range values {
		minH.Insert(v)
	}
	peek, _ := minH.Peek()
	fmt.Printf("  MinHeap — internal slice: %v\n", minH.Snapshot())
	fmt.Printf("  MinHeap — Peek (min):     %d\n", peek)
	fmt.Print("  MinHeap — Extract order:  [")
	first := true
	for !minH.IsEmpty() {
		v, _ := minH.ExtractMin()
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(v)
		first = false
	}
	fmt.Println("]  ← ascending order")

	fmt.Println()

	// MaxHeap
	maxH := heap.NewMaxHeap()
	for _, v := range values {
		maxH.Insert(v)
	}
	peekMax, _ := maxH.Peek()
	fmt.Printf("  MaxHeap — internal slice: %v\n", maxH.Snapshot())
	fmt.Printf("  MaxHeap — Peek (max):     %d\n", peekMax)
	fmt.Print("  MaxHeap — Extract order:  [")
	first = true
	for !maxH.IsEmpty() {
		v, _ := maxH.ExtractMax()
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(v)
		first = false
	}
	fmt.Println("]  ← descending order")
}

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

func section(title string) {
	fmt.Println()
	fmt.Printf("━━━  %s  ━━━\n\n", title)
}
