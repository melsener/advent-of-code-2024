package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strconv"
	"strings"
)

var END_POS = 71

func init() {
	common.RegisterSolver(18, SolveDay18)
}

func shortestPath(grid [][]rune, start []int, end []int) int {
	var dirs = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	if grid[start[0]][start[1]] == '#' || grid[end[0]][end[1]] == '#' {
		return -1
	}

	queue := [][]int{{start[0], start[1], 0}}
	key := strconv.Itoa(start[0]) + ":" + strconv.Itoa(start[1])
	visited := map[string]bool{key: true}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item[0] == end[0] && item[1] == end[1] {
			return item[2]
		}

		for i := range dirs {
			r := item[0] + dirs[i][0]
			c := item[1] + dirs[i][1]

			key := strconv.Itoa(r) + ":" + strconv.Itoa(c)
			_, ok := visited[key]
			if ok {
				continue
			}
			if r >= 0 && r < END_POS && c >= 0 && c < END_POS && grid[r][c] != '#' {
				queue = append(queue, []int{r, c, item[2] + 1})
				visited[key] = true
			}
		}
	}

	return -1
}

func SolveDay18(input string) string {
	lines := strings.Split(input, "\n")
	n := len(lines)
	bytePositions := make([][]int, n)
	for i := range lines {
		coordinates := strings.Split(lines[i], ",")
		x, err1 := strconv.Atoi(coordinates[0])
		y, err2 := strconv.Atoi(coordinates[1])
		if err1 == nil && err2 == nil {
			bytePositions[i] = make([]int, 2)
			bytePositions[i][0] = y
			bytePositions[i][1] = x
		}
	}

	grid := make([][]rune, END_POS)
	for i := range grid {
		grid[i] = []rune(".......................................................................")
	}

	for i := 0; i < 1024; i++ {
		pos := bytePositions[i]
		grid[pos[0]][pos[1]] = '#'
	}

	// for i := range grid {
	// 	for j := range grid[i] {
	// 		fmt.Print(string(grid[i][j]))
	// 	}
	// 	fmt.Print("\n")
	// }

	result1 := shortestPath(grid, []int{0, 0}, []int{END_POS - 1, END_POS - 1})
	result2 := 0
	// 310 0
	return fmt.Sprintf("%d %d", result1, result2)
}
