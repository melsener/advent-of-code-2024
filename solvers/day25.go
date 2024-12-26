package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strings"
)

func init() {
	common.RegisterSolver(25, SolveDay25)
}

func processLock(lock []string) []int {
	lockGrid := [][]rune{}
	for i := range lock {
		line := []rune(lock[i])
		lockGrid = append(lockGrid, line)
	}
	nums := []int{}
	for i := 0; i < 5; i++ {
		res := 0
		for j := 1; j < len(lock); j++ {
			if lockGrid[j][i] == '#' {
				res += 1
			}
		}
		nums = append(nums, res)
	}
	return nums
}

func processKey(key []string) []int {
	keyGrid := [][]rune{}
	for i := range key {
		line := []rune(key[i])
		keyGrid = append(keyGrid, line)
	}
	nums := []int{}
	for i := 0; i < 5; i++ {
		res := 0
		for j := len(key) - 2; j >= 0; j-- {
			if keyGrid[j][i] == '#' {
				res += 1
			}
		}
		nums = append(nums, res)
	}
	return nums
}

func isOverlapping(lock []int, key []int) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return true
		}
	}
	return false
}

func SolveDay25(input string) string {
	schematics := strings.Split(input, "\n\n")
	locks := [][]int{}
	keys := [][]int{}
	for i := range schematics {
		lines := strings.Split(schematics[i], "\n")
		// lock
		if lines[0] == "#####" {
			lock := processLock(lines)
			locks = append(locks, lock)
		} else {
			// key
			key := processKey(lines)
			keys = append(keys, key)
		}
	}
	result1 := 0
	for _, key := range keys {
		for _, lock := range locks {
			if !isOverlapping(lock, key) {
				result1 += 1
			}
		}
	}
	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
