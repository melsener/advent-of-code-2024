package solvers

import (
	"aoc-2024/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(3, SolveDay3)
}

func SolveDay3(input string) string {
	lines := strings.Split(input, "\n")
	input = strings.Join(lines, "\\n") // Trick

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	result1 := 0
	for _, match := range matches {
		x, err1 := strconv.Atoi(match[1])
		y, err2 := strconv.Atoi(match[2])

		if err1 == nil && err2 == nil {
			result1 += (x * y)
		} else {
			fmt.Println("Error parsing numbers")
		}
	}

	re2 := regexp.MustCompile(`don't\(\).*?do\(\)`)
	modified := re2.ReplaceAllString(input, "")
	re3 := regexp.MustCompile(`don't\(\).*$`)
	modified = re3.ReplaceAllString(modified, "")

	result2 := 0
	matches = re.FindAllStringSubmatch(modified, -1)
	for _, match := range matches {
		x, err1 := strconv.Atoi(match[1])
		y, err2 := strconv.Atoi(match[2])

		if err1 == nil && err2 == nil {
			result2 += (x * y)
		} else {
			fmt.Println("Error parsing numbers")
		}
	}

	// 171183089 63866497
	return fmt.Sprintf("%d %d", result1, result2)
}
