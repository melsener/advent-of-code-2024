package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strings"
)

func init() {
	common.RegisterSolver(2, SolveDay2)
}

func isValidRow(row []string, i int) bool {
	prev := convertNum(row[0], i)
	var sign int = 0
	for j := 1; j < len(row); j++ {
		cur := convertNum(row[j], i)

		if prev == cur {
			return false
		}

		diff := prev - cur

		if diff < 0 && sign == -1 {
			// Increasing
			return false
		} else if diff > 0 && sign == 1 {
			// Decreasing
			return false
		}

		if sign == 0 {
			if diff < 0 {
				sign = 1
			} else {
				sign = -1
			}
		}

		if diff < 0 {
			diff *= -1
		}

		if diff > 3 {
			return false
		}

		prev = cur
	}

	return true
}

func SolveDay2(input string) string {
	lines := strings.Split(input, "\n")
	n := len(lines)

	var unsafe int = 0
	for i := 0; i < n; i++ {
		reports := strings.Fields(lines[i])

		res := isValidRow(reports, i)
		if !res {
			unsafe += 1
		}
	}
	result1 := n - unsafe
	// Part 2
	result2 := 0
	for i := 0; i < n; i++ {
		reports := strings.Fields(lines[i])
		res := isValidRow(reports, i)
		if res {
			result2 += 1
		} else {
			for j := 0; j < len(reports); j++ {
				slice := append([]string{}, reports[:j]...)
				slice = append(slice, reports[j+1:]...)
				res := isValidRow(slice, i)
				if res {
					result2 += 1
					break
				}
			}
		}
	}

	return fmt.Sprintf("%d %d", result1, result2)
}
