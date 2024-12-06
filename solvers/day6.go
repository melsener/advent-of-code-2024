package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strings"
)

func init() {
	common.RegisterSolver(6, SolveDay6)
}

func step(curPos []int, dir rune) []int {
	r := curPos[0]
	c := curPos[1]

	switch dir {
	case 'U':
		return []int{r - 1, c}
	case 'D':
		return []int{r + 1, c}
	case 'L':
		return []int{r, c - 1}
	case 'R':
		return []int{r, c + 1}
	}
	return []int{r, c}
}

func turnRight(dir rune) rune {
	directions := []rune{'U', 'R', 'D', 'L'}
	var result int
	for i := range directions {
		if directions[i] == dir {
			result = i
			break
		}
	}

	if result == 3 {
		return directions[0]
	}
	return directions[result+1]
}

func isInside(curPos []int, n int, m int) bool {
	r := curPos[0]
	c := curPos[1]

	return (r > -1) && (r < n) && (c > -1) && (c < m)
}

func SolveDay6(input string) string {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	grid := make([][]rune, n)
	visited := make([][]int, n)
	currentPos := make([]int, 2)
	dir := 'O'
	for i := range grid {
		visited[i] = make([]int, m)
		grid[i] = []rune(lines[i])
		up := strings.IndexRune(lines[i], '^')
		down := strings.IndexRune(lines[i], 'v')
		left := strings.IndexRune(lines[i], '<')
		right := strings.IndexRune(lines[i], '>')
		if up > -1 {
			currentPos[0] = i
			currentPos[1] = up
			dir = 'U'
		} else if down > -1 {
			currentPos[0] = i
			currentPos[1] = down
			dir = 'D'
		} else if left > -1 {
			currentPos[0] = i
			currentPos[1] = left
			dir = 'L'
		} else if right > -1 {
			currentPos[0] = i
			currentPos[1] = right
			dir = 'R'
		}
	}
	visited[currentPos[0]][currentPos[1]] = 1
	inside := true

	for inside {
		next := step(currentPos, dir)
		nextInside := isInside(next, n, m)
		if !nextInside {
			break
		}
		if string(grid[next[0]][next[1]]) == "#" {
			dir = turnRight(dir)
		} else {
			currentPos[0] = next[0]
			currentPos[1] = next[1]
			visited[currentPos[0]][currentPos[1]] = 1
		}

		inside = isInside(currentPos, n, m)
	}

	result1 := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == 1 {
				result1 += 1
			}
		}
	}

	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
