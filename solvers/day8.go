package solvers

import (
	"aoc-2024/common"
	"fmt"
	"regexp"
	"strings"
)

func init() {
	common.RegisterSolver(8, SolveDay8)
}

func antennaDif(f [2]int, s [2]int) [2]int {
	return [2]int{f[0] - s[0], f[1] - s[1]}
}

func antinodeWithinLimit(coordinates [2]int, limitY int, limitX int) bool {
	r := coordinates[0]
	c := coordinates[1]

	return (r > -1) && (r < limitY) && (c > -1) && (c < limitX)
}

func countAntinodes(coordinates [][2]int, limitY int, limitX int, visited [][]rune) {
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			f := coordinates[i]
			s := coordinates[j]
			diff := antennaDif(f, s)

			p11 := f[0] + diff[0]
			p12 := f[1] + diff[1]
			if antinodeWithinLimit([2]int{p11, p12}, limitY, limitX) {
				visited[p11][p12] = rune('#')
			}
			p21 := s[0] - diff[0]
			p22 := s[1] - diff[1]
			if antinodeWithinLimit([2]int{p21, p22}, limitY, limitX) {
				visited[p21][p22] = rune('#')
			}
		}
	}
}

func SolveDay8(input string) string {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`\.`)
	antennaCoordinates := make(map[rune][][2]int)
	for i := range lines {
		res := re.ReplaceAllString(lines[i], "")
		resSplit := strings.Split(res, "")
		if len(resSplit) == 0 {
			continue
		}
		cur := lines[i]
		for j := range cur {
			if cur[j] == '.' {
				continue
			}
			char := rune(cur[j])
			value, exists := antennaCoordinates[char]
			if !exists {
				antennaCoordinates[char] = [][2]int{}
			}
			antennaCoordinates[char] = append(value, [2]int{i, j})
		}
	}

	visited := make([][]rune, len(lines))
	for k := range visited {
		visited[k] = make([]rune, len(lines[0]))
	}

	result1 := 0
	for _, value := range antennaCoordinates {
		if len(value) < 2 {
			continue
		}
		countAntinodes(value, len(lines), len(lines[0]), visited)
	}

	for i := range visited {
		for j := range visited[i] {
			if string(visited[i][j]) == "#" {
				result1 += 1
			}
		}
	}

	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
