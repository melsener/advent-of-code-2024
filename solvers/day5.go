package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(5, SolveDay5)
}

func search(row []int, num int) bool {
	for i := range row {
		if row[i] == num {
			return true
		}
	}

	return false
}

func isValid(row []int, sources map[int][]int, targets map[int][]int) bool {
	valid := true
	prev := -1
	var index int
	for i := range row {
		_, ok := sources[row[i]]
		if ok {
			prev = row[i]
			index = i
			break
		}
	}

	if prev != -1 {
		for i := index + 1; i < len(row); i++ {
			_, ok1 := targets[row[i]]
			_, ok2 := sources[row[i]]

			if !ok1 && !ok2 {
				continue
			}

			if ok1 {
				found := search(targets[row[i]], prev)
				if !found {
					valid = false
					break
				}
				prev = row[i]
			} else if ok2 {
				found := search(sources[prev], row[i])
				if !found {
					valid = false
					break
				}
			}

		}
	}

	return valid
}

func SolveDay5(input string) string {
	parts := strings.Split(input, "\n\n")

	rules := strings.Split(parts[0], "\n")
	updates := strings.Split(parts[1], "\n")

	sources := make(map[int][]int)
	targets := make(map[int][]int)
	for i := range rules {
		nums := strings.Split(rules[i], "|")
		source, err1 := strconv.Atoi(nums[0])
		target, err2 := strconv.Atoi(nums[1])
		if err1 == nil && err2 == nil {
			_, ok := sources[source]
			if !ok {
				sources[source] = make([]int, 0, 100)
			}
			sources[source] = append(sources[source], target)

			_, ok2 := targets[target]
			if !ok2 {
				targets[target] = make([]int, 0, 100)
			}
			targets[target] = append(targets[target], source)
		} else {
			fmt.Println("Error parsing numbers")
		}
	}

	result1 := 0
	for i := range updates {
		nn := strings.Split(updates[i], ",")
		nums := make([]int, 0, len(nn))
		for j := range nn {
			num, err := strconv.Atoi(nn[j])
			if err == nil {
				nums = append(nums, num)
			}
		}

		valid := isValid(nums, sources, targets)
		if valid {
			fmt.Print("valid", nums, "\n")
			mid := len(nums) / 2
			result1 += nums[mid]
		}
	}

	fmt.Print(sources, "\n")
	fmt.Print(targets, "\n")

	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
