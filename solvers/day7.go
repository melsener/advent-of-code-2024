package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(7, SolveDay7)
}

func isValidEquation(result int, nums []int, currentVal int, currentIndex int) bool {
	if currentIndex == len(nums) {
		return currentVal == result
	}

	for i := currentIndex; i < len(nums); i++ {
		isSum := isValidEquation(result, nums, currentVal+nums[i], i+1)
		if isSum {
			return true
		}

		isProd := isValidEquation(result, nums, currentVal*nums[i], i+1)
		if isProd {
			return true
		}
	}
	return false
}

func SolveDay7(input string) string {
	lines := strings.Split(input, "\n")
	result1 := 0
	result2 := 0

	for i := range lines {
		equation := strings.Split(lines[i], ":")
		res, err1 := strconv.Atoi(equation[0])
		if err1 == nil {
			nn := strings.Fields(equation[1])
			nums := make([]int, 0, len(nn))
			for j := range nn {
				num, err := strconv.Atoi(nn[j])
				if err == nil {
					nums = append(nums, num)
				}
			}
			valid := isValidEquation(res, nums, nums[0], 1)
			if valid {
				result1 += res
			}
		}
	}
	// 1985268524462 0
	return fmt.Sprintf("%d %d", result1, result2)
}
