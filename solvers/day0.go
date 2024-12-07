package solvers

import (
	"aoc-2024/common"
	"fmt"
)

func init() {
	common.RegisterSolver(0, SolveDay0)
}

func SolveDay0(input string) string {
	result1 := 0
	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
