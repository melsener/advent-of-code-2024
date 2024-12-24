package solvers

import (
	"aoc-2024/common"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(24, SolveDay24)
}

func calculate(a bool, b bool, op string) bool {
	switch op {
	case "AND":
		return a && b
	case "OR":
		return a || b
	case "XOR":
		return (!a && b) || (a && !b)
	}
	return false
}

func SolveDay24(input string) string {
	inputs := strings.Split(input, "\n\n")
	inp1 := strings.Split(inputs[0], "\n")
	inp2 := strings.Split(inputs[1], "\n")
	wires := make(map[string]bool)
	for i := range inp1 {
		wire := strings.Split(inp1[i], ": ")
		if wire[1] == "0" {
			wires[wire[0]] = false
		} else {
			wires[wire[0]] = true
		}
	}

	queue := [][]string{}
	for i := range inp2 {
		parts := strings.Split(inp2[i], " -> ")
		re := regexp.MustCompile(`(\w+)\s+(AND|XOR|OR)\s+(\w+)\s+->\s+(\w+)`)
		matches := re.FindStringSubmatch(inp2[i])
		if len(matches) > 0 {
			firstOp := matches[1]
			secondOp := matches[3]
			operation := matches[2]

			f, ok1 := wires[firstOp]
			s, ok2 := wires[secondOp]
			if !ok1 || !ok2 {
				queue = append(queue, []string{firstOp, secondOp, operation, parts[1]})
			} else {
				res := calculate(f, s, operation)
				wires[parts[1]] = res
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		firstOp := cur[0]
		secondOp := cur[1]
		operation := cur[2]
		outcome := cur[3]

		f, ok1 := wires[firstOp]
		s, ok2 := wires[secondOp]
		if !ok1 || !ok2 {
			queue = append(queue, []string{firstOp, secondOp, operation, outcome})
		} else {
			res := calculate(f, s, operation)
			wires[outcome] = res
		}
	}

	zCount := 0
	for k := range wires {
		if strings.HasPrefix(k, "z") {
			zCount++
		}
	}

	result1 := 0
	for i := zCount; i >= 0; i-- {
		key := ""
		if i > 9 {
			key = "z" + strconv.Itoa(i)
		} else {
			key = "z0" + strconv.Itoa(i)
		}

		res := wires[key]
		val := 0
		if res {
			val = 1
		}
		result1 += (val * int(math.Pow(2, float64(i))))
	}
	result2 := 0
	// 60714423975686 0
	return fmt.Sprintf("%d %d", result1, result2)
}
