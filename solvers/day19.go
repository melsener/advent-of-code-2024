package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strings"
)

func init() {
	common.RegisterSolver(19, SolveDay19)
}

func possibleDesign(design string, patterns []string, cur string) bool {
	if cur == design {
		return true
	} else if len(cur) > len(design) {
		return false
	} else if cur != design[:len(cur)] {
		return false
	}

	for i := range patterns {
		res := possibleDesign(design, patterns, cur+patterns[i])
		if res {
			return true
		}
	}
	return false
}

func SolveDay19(input string) string {
	inputs := strings.Split(input, "\n\n")
	patterns := strings.Split(inputs[0], ", ")
	designs := strings.Split(inputs[1], "\n")

	result1 := 0
	for i := range designs {
		res := possibleDesign(designs[i], patterns, "")
		if res {
			result1 += 1
		}
	}
	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
