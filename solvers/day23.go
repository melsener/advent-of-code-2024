package solvers

import (
	"aoc-2024/common"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func init() {
	common.RegisterSolver(23, SolveDay23)
}

func findCommonElements(c1 string, c2 string, connections map[string]map[string]int) []string {
	results := []string{}
	for k := range connections[c1] {
		if k != c2 {
			_, ok := connections[c2][k]
			if ok {
				set := []string{c1, c2, k}
				sort.Strings(set)
				results = append(results, strings.Join(set, "-"))
			}
		}
	}

	return results
}

func SolveDay23(input string) string {
	lines := strings.Split(input, "\n")
	connections := map[string]map[string]int{}
	for _, line := range lines {
		computers := strings.Split(line, "-")
		_, ok := connections[computers[0]]
		if !ok {
			connections[computers[0]] = map[string]int{}
		}
		connections[computers[0]][computers[1]] = 1
		// --
		_, ok = connections[computers[1]]
		if !ok {
			connections[computers[1]] = map[string]int{}
		}
		connections[computers[1]][computers[0]] = 1
	}

	triplets := map[string]bool{}
	for key, val := range connections {
		for k := range val {
			res := findCommonElements(key, k, connections)
			for i := range res {
				triplets[res[i]] = true
			}
		}
	}

	result1 := 0
	for k := range triplets {
		re := regexp.MustCompile(`(^|-)t`)
		historian := re.MatchString(k)
		if historian {
			result1 += 1
		}
	}

	result2 := 0
	// 926 0
	return fmt.Sprintf("%d %d", result1, result2)
}
