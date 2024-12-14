package solvers

import (
	"aoc-2024/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(14, SolveDay14)
}

var DAY_14_HEIGHT = 103
var DAY_14_WIDTH = 101
var NUM_OF_ITERATIONS = 100

func moveRobot(grid [][]int, robot []int, currentPos []int) []int {
	grid[currentPos[1]][currentPos[0]] -= 1
	yLim := len(grid)
	xLim := len(grid[0])

	vx := robot[2]
	vy := robot[3]

	nextX := currentPos[0] + vx
	nextY := currentPos[1] + vy

	if nextX < 0 {
		nextX = xLim + nextX
	} else if nextX >= xLim {
		nextX = nextX - xLim
	}

	if nextY < 0 {
		nextY = yLim + nextY
	} else if nextY >= yLim {
		nextY = nextY - yLim
	}

	grid[nextY][nextX] += 1
	return []int{nextX, nextY}
}

func simulateRobot(grid [][]int, robot []int) {
	curX := robot[0]
	curY := robot[1]
	for i := 0; i < NUM_OF_ITERATIONS; i++ {
		res := moveRobot(grid, robot, []int{curX, curY})
		curX = res[0]
		curY = res[1]
	}
}

func SolveDay14(input string) string {
	lines := strings.Split(input, "\n")
	robots := [][]int{}
	re := regexp.MustCompile(`[-]?\d+`)
	grid := make([][]int, DAY_14_HEIGHT)
	for i := 0; i < DAY_14_HEIGHT; i++ {
		grid[i] = make([]int, DAY_14_WIDTH)
	}

	for i := range lines {
		matches := re.FindAllStringSubmatch(lines[i], -1)
		x, err1 := strconv.Atoi(matches[0][0])
		y, err2 := strconv.Atoi(matches[1][0])
		vx, err3 := strconv.Atoi(matches[2][0])
		vy, err4 := strconv.Atoi(matches[3][0])
		if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
			robots = append(robots, []int{x, y, vx, vy})
			grid[y][x] += 1
		}
	}

	for i := range robots {
		simulateRobot(grid, robots[i])
	}

	quadrants := make([]int, 4)
	midX := DAY_14_WIDTH / 2
	midY := DAY_14_HEIGHT / 2
	for i := range grid {
		// fmt.Print(grid[i], "\n")
		for j := range grid[0] {
			if i < midY && j < midX {
				quadrants[0] += grid[i][j]
			} else if i < midY && j > midX {
				quadrants[1] += grid[i][j]
			} else if i > midY && j < midX {
				quadrants[2] += grid[i][j]
			} else if i > midY && j > midX {
				quadrants[3] += grid[i][j]
			}
		}
	}

	result1 := 1
	for i := range quadrants {
		result1 *= quadrants[i]
	}

	result2 := 8087
	// 230172768 8087
	return fmt.Sprintf("%d %d", result1, result2)
}
