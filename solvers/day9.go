package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(9, SolveDay9)
}

func createArray(value int, size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = value
	}
	return arr
}

func SolveDay9(input string) string {
	nums := strings.Split(input, "")
	fsys := []int{}

	fileId := 0
	for i := range nums {
		num, err := strconv.Atoi(nums[i])
		if err != nil {
			break
		}
		if i%2 == 0 {
			fsys = append(fsys, createArray(fileId, num)...)
			fileId += 1
		} else {
			fsys = append(fsys, createArray(-1, num)...)
		}
	}

	j := len(fsys) - 1
	for i := 0; i < len(fsys); i++ {
		for fsys[j] == -1 {
			j--
		}

		if j < i {
			break
		}

		if fsys[i] == -1 {
			fsys[i] = fsys[j]
			fsys[j] = -1
		}
	}

	result1 := 0
	for i := 0; i < len(fsys); i++ {
		if fsys[i] == -1 {
			break
		}
		result1 += (i * fsys[i])
	}

	result2 := 0
	// 6390180901651
	return fmt.Sprintf("%d %d", result1, result2)
}
