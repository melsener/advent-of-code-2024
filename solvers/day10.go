package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(10, SolveDay10)
}

func hike(trailHead []int, trail [][]int, visited [][]bool) int {
	score := 0
	r := trailHead[0]
	c := trailHead[1]
	if r >= len(trail) || c >= len(trail[0]) {
		return 0
	}

	visited[r][c] = true
	if trail[r][c] == 9 {
		return 1
	}

	var dirs = [][]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}

	for i := range dirs {
		rNext := r + dirs[i][0]
		cNext := c + dirs[i][1]

		if (rNext > -1) && (rNext < len(trail)) && (cNext > -1) && (cNext < len(trail[0])) {
			if !visited[rNext][cNext] && (trail[rNext][cNext] == trail[r][c]+1) {
				score += hike([]int{rNext, cNext}, trail, visited)
			}
		}
	}

	return score
}

func SolveDay10(input string) string {
	lines := strings.Split(input, "\n")
	trail := [][]int{}
	trailHeads := [][]int{}
	for i := range lines {
		nn := strings.Split(lines[i], "")
		trail = append(trail, []int{})
		for j := range nn {
			num, err := strconv.Atoi(nn[j])
			if err == nil {
				trail[i] = append(trail[i], num)
				if num == 0 {
					trailHeads = append(trailHeads, []int{i, j})
				}
			}
		}
	}

	result1 := 0
	for i := range trailHeads {
		visited := make([][]bool, len(lines))
		for k := range visited {
			visited[k] = make([]bool, len(lines[0]))
		}
		result1 += hike(trailHeads[i], trail, visited)
	}

	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
