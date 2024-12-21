package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strings"
)

func init() {
	common.RegisterSolver(20, SolveDay20)
}

func raceCondition(maze [][]rune, n int, m int, start []int, end []int) int {
	visited := make([][]bool, n)
	for k := range visited {
		visited[k] = make([]bool, m)
	}

	var dirs = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	queue := [][]int{{start[0], start[1], 0}}
	visited[start[0]][start[1]] = true

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state[0] == end[0] && state[1] == end[1] {
			return state[2]
		}

		for i := range dirs {
			nY := state[0] + dirs[i][0]
			nX := state[1] + dirs[i][1]

			if (nY > -1) && (nY < n) && (nX > -1) && (nX < m) && maze[nY][nX] != '#' && !visited[nY][nX] {
				queue = append(queue, []int{nY, nX, state[2] + 1})
				visited[nY][nX] = true
			}
		}
	}

	return -1
}

func copyGrid(grid [][]rune, n int, m int) [][]rune {
	copy := make([][]rune, n)
	for i := range grid {
		copy[i] = make([]rune, m)
		for j := range grid[i] {
			copy[i][j] = grid[i][j]
		}
	}
	return copy
}

func SolveDay20(input string) string {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	startPos := make([]int, 2)
	endPos := make([]int, 2)
	maze := make([][]rune, n)
	for i := range lines {
		maze[i] = []rune(lines[i])
		s := strings.IndexRune(lines[i], 'S')
		e := strings.IndexRune(lines[i], 'E')
		if s > -1 {
			startPos[0] = i
			startPos[1] = s
		}
		if e > -1 {
			endPos[0] = i
			endPos[1] = e
		}
	}

	actual := raceCondition(maze, n, m, startPos, endPos)
	countMap := make(map[int]int)
	result1 := 0
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == '#' {
				modifiedMaze := copyGrid(maze, n, m)
				modifiedMaze[i][j] = '.'
				res := raceCondition(modifiedMaze, n, m, startPos, endPos)
				if res > -1 && res < actual {
					diff := actual - res
					_, ok := countMap[diff]
					if !ok {
						countMap[diff] = 0
					}
					countMap[diff] += 1

					if diff >= 100 {
						result1 += 1
					}
				}
			}
		}
	}

	// fmt.Print(actual, "\n")
	// fmt.Print(countMap, "\n")

	result2 := 0
	// 1358 0 very bad perf
	return fmt.Sprintf("%d %d", result1, result2)
}
