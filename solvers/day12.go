package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strings"
)

func init() {
	common.RegisterSolver(12, SolveDay12)
}

func calculateAreaPerimeter(r int, c int, grid [][]rune, visited [][]bool) []int {
	var dirs = [][]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}

	visited[r][c] = true
	v0 := 1
	v1 := 4
	for i := range dirs {
		rNext := r + dirs[i][0]
		cNext := c + dirs[i][1]

		if (rNext > -1) && (rNext < len(grid)) && (cNext > -1) && (cNext < len(grid[0])) {
			if grid[rNext][cNext] == grid[r][c] {
				if !visited[rNext][cNext] {
					res := calculateAreaPerimeter(rNext, cNext, grid, visited)
					v0 += res[0]
					v1 += res[1]
				}
				v1--
			}
		}
	}

	return []int{v0, v1}
}

func SolveDay12(input string) string {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	grid := make([][]rune, n)
	visited := make([][]bool, n)
	for i := range lines {
		grid[i] = []rune(lines[i])
		visited[i] = make([]bool, m)
	}

	result1 := 0

	for i := range grid {
		for j := range grid[i] {
			if !visited[i][j] {
				res := calculateAreaPerimeter(i, j, grid, visited)
				result1 += (res[0] * res[1])
			}
		}
	}

	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
