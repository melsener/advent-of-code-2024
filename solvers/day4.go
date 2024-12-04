package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strings"
)

func init() {
	common.RegisterSolver(4, SolveDay4)
}

func checkDirection(word string, dir []int, n int, m int, grid [][]string, i int, j int) bool {
	for l := 0; l < len(word); l++ {
		row := i + l*dir[0]
		col := j + l*dir[1]
		if row < 0 || col < 0 || row >= n || col >= m {
			return false
		} else if grid[row][col] != string(word[l]) {
			return false
		}
	}
	return true
}

func SolveDay4(input string) string {
	word := "XMAS"
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	grid := make([][]string, n)
	for i := range grid {
		grid[i] = strings.Split(lines[i], "")
	}

	visited := make([][]bool, n)
	for k := range visited {
		visited[k] = make([]bool, m)
	}

	result1 := 0

	var checks = [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < len(checks); k++ {
				dir := checkDirection(word, checks[k], n, m, grid, i, j)

				if dir {
					result1 += 1
				}
			}
		}
	}

	result2 := 0
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			marksTheSpot := 0
			if grid[i][j] == "A" {
				topLeft := grid[i-1][j-1]
				topRight := grid[i-1][j+1]
				bottomLeft := grid[i+1][j-1]
				bottomRight := grid[i+1][j+1]

				// \ part
				if (topLeft == "M" && bottomRight == "S") || (topLeft == "S" && bottomRight == "M") {
					marksTheSpot += 1
				}

				// / part
				if (topRight == "M" && bottomLeft == "S") || (topRight == "S" && bottomLeft == "M") {
					marksTheSpot += 1
				}
			}

			if marksTheSpot == 2 {
				result2 += 1
			}
		}
	}

	return fmt.Sprintf("%d %d", result1, result2)
}
