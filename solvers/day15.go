package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strings"
)

func init() {
	common.RegisterSolver(15, SolveDay15)
}

func moveBall(grid [][]rune, n int, m int, currentPos []int, nextPos []int, currentDir rune) []int {
	numOfWalls := 0
	free := -1
	if currentDir == '^' {
		for i := nextPos[0]; i > -1; i-- {
			if grid[i][nextPos[1]] == '.' {
				free = i
				break
			} else if grid[i][nextPos[1]] == 'O' {
				numOfWalls++
			} else {
				return currentPos
			}
		}

		if free > -1 {
			grid[currentPos[0]][currentPos[1]] = '.'
			grid[free][nextPos[1]] = 'O'
			grid[nextPos[0]][nextPos[1]] = '@'
			return nextPos
		}
	} else if currentDir == 'v' {
		for i := nextPos[0]; i < n; i++ {
			if grid[i][nextPos[1]] == '.' {
				free = i
				break
			} else if grid[i][nextPos[1]] == 'O' {
				numOfWalls++
			} else {
				return currentPos
			}
		}
		if free > -1 {
			grid[currentPos[0]][currentPos[1]] = '.'
			grid[free][nextPos[1]] = 'O'
			grid[nextPos[0]][nextPos[1]] = '@'
			return nextPos
		}
	} else if currentDir == '<' {
		for i := nextPos[1]; i > -1; i-- {
			if grid[nextPos[0]][i] == '.' {
				free = i
				break
			} else if grid[nextPos[0]][i] == 'O' {
				numOfWalls++
			} else {
				return currentPos
			}
		}
		if free > -1 {
			grid[currentPos[0]][currentPos[1]] = '.'
			grid[nextPos[0]][free] = 'O'
			grid[nextPos[0]][nextPos[1]] = '@'
			return nextPos
		}
	} else {
		for i := nextPos[1]; i < m; i++ {
			if grid[nextPos[0]][i] == '.' {
				free = i
				break
			} else if grid[nextPos[0]][i] == 'O' {
				numOfWalls++
			} else {
				return currentPos
			}
		}
		if free > -1 {
			grid[currentPos[0]][currentPos[1]] = '.'
			grid[nextPos[0]][free] = 'O'
			grid[nextPos[0]][nextPos[1]] = '@'
			return nextPos
		}
	}
	return currentPos
}

func moveRobot15(grid [][]rune, n int, m int, currentPos []int, currentDir rune) []int {
	nextPos := []int{currentPos[0], currentPos[1]}
	if currentDir == '^' {
		nextPos[0] -= 1
	} else if currentDir == 'v' {
		nextPos[0] += 1
	} else if currentDir == '<' {
		nextPos[1] -= 1
	} else {
		nextPos[1] += 1
	}

	if grid[nextPos[0]][nextPos[1]] == '.' {
		grid[nextPos[0]][nextPos[1]] = '@'
		grid[currentPos[0]][currentPos[1]] = '.'
		return nextPos
	} else if grid[nextPos[0]][nextPos[1]] == '#' {
		return currentPos
	}
	return moveBall(grid, n, m, currentPos, nextPos, currentDir)
}

func SolveDay15(input string) string {
	inputs := strings.Split(input, "\n\n")

	gridLines := strings.Split(inputs[0], "\n")
	n := len(gridLines)
	m := len(gridLines[0])

	grid := make([][]rune, n)
	curPos := make([]int, 2)
	for i := range gridLines {
		grid[i] = []rune(gridLines[i])

		robot := strings.IndexRune(gridLines[i], '@')
		if robot > -1 {
			curPos[0] = i
			curPos[1] = robot
		}
	}

	moves := []rune{}
	moveLines := strings.Split(inputs[1], "\n")
	for i := range moveLines {
		chars := []rune(moveLines[i])
		moves = append(moves, chars...)
	}

	for i := range moves {
		curPos = moveRobot15(grid, n, m, curPos, moves[i])
	}

	result1 := 0
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(string(grid[i][j]))
			if grid[i][j] == 'O' {
				result1 += (100*i + j)
			}
		}
		fmt.Print("\n")
	}

	result2 := 0
	// 1398947 0
	return fmt.Sprintf("%d %d", result1, result2)
}
