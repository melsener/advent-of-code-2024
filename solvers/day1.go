package solvers

import (
	"aoc-2024/common"
	"container/heap"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(1, SolveDay1)
}

// BEGIN: Taken from example here https://pkg.go.dev/container/heap
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// END:

func convertNum(input string, lineNum int) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Invalid num1: %v on line %d", input, lineNum)
	}
	return num
}

func SolveDay1(input string) string {
	lines := strings.Split(input, "\n")
	n := len(lines)
	left := &IntHeap{}
	heap.Init(left)

	right := &IntHeap{}
	heap.Init(right)

	for i := 0; i < n; i++ {
		words := strings.Fields(lines[i])
		// First num
		num1 := convertNum(words[0], i)
		heap.Push(left, num1)
		// Second num
		num2 := convertNum(words[1], i)
		heap.Push(right, num2)
	}

	var result1 int = 0
	for i := 0; i < n; i++ {
		lPop := heap.Pop(left)
		rPop := heap.Pop(right)
		l := lPop.(int)
		r := rPop.(int)

		if l < r {
			result1 += (r - l)
		} else {
			result1 += (l - r)
		}
	}

	// Part 2
	similarityMap := make(map[int]int)
	// Add map keys
	for i := 0; i < n; i++ {
		words := strings.Fields(lines[i])
		// First num
		num1 := convertNum(words[0], i)
		similarityMap[num1] = 0
	}
	// Count frequencies
	for i := 0; i < n; i++ {
		words := strings.Fields(lines[i])
		// Second num
		num2 := convertNum(words[1], i)
		val, ok := similarityMap[num2]
		if ok {
			similarityMap[num2] = val + 1
		}
	}

	// Calculate the result
	var result2 int = 0
	for i := 0; i < n; i++ {
		words := strings.Fields(lines[i])
		// First num
		num1 := convertNum(words[0], i)
		val, ok := similarityMap[num1]
		if ok {
			result2 += (num1 * val)
		}
	}

	// 2000468 18567089
	return fmt.Sprintf("%d %d", result1, result2)
}
