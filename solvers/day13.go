package solvers

import (
	"aoc-2024/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(13, SolveDay13)
}

func extractInts(input string, re *regexp.Regexp) []int {
	matches := re.FindAllString(input, -1)
	ints := make([]int, len(matches))
	for i, match := range matches {
		val, err := strconv.Atoi(match)
		if err != nil {
			fmt.Printf("Error converting %s to integer: %v\n", match, err)
			continue
		}
		ints[i] = val
	}
	return ints
}

func calculateCost(a []int, b []int, prize []int) int {
	MaxUint := ^uint(0)
	min := int(MaxUint >> 1)
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			x := (i * a[0]) + (j * b[0])
			y := (i * a[1]) + (j * b[1])

			if x == prize[0] && y == prize[1] {
				cost := (i * 3) + j
				if cost < min {
					min = cost
				}
			}
		}
	}

	if min == int(MaxUint>>1) {
		return 0
	}
	return min
}

func SolveDay13(input string) string {
	lines := strings.Split(input, "\n\n")
	n := len(lines)
	re := regexp.MustCompile(`[-]?\d+`)
	arr := [][][]int{}
	for i := range lines {
		rules := strings.Split(lines[i], "\n")
		a := extractInts(rules[0], re)
		b := extractInts(rules[1], re)
		prize := extractInts(rules[2], re)
		res := [][]int{a, b, prize}
		arr = append(arr, res)
	}

	result1 := 0
	for i := 0; i < n; i++ {
		res := calculateCost(arr[i][0], arr[i][1], arr[i][2])
		result1 += res
	}

	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
