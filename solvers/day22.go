package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strconv"
	"strings"
)

var LIMIT = 2000

func init() {
	common.RegisterSolver(22, SolveDay22)
}

func nextSecret(num int) int {
	res := num * 64
	num = res ^ num
	num = num % 16777216
	// --
	x := int(num / 32)
	num = x ^ num
	num = num % 16777216
	// --
	x = num * 2048
	num = x ^ num
	num = num % 16777216
	return num
}

func calculateSecret(num int, globalCache map[string]int) int {
	cur := num
	diff := []int{}
	cache := make(map[string]int)
	for i := 0; i < LIMIT; i++ {
		next := nextSecret(cur)
		diff = append(diff, (next%10)-(cur%10))
		cur = next

		if len(diff) == 4 {
			key := strconv.Itoa(diff[0]) + ":" + strconv.Itoa(diff[1]) + ":" + strconv.Itoa(diff[2]) + ":" + strconv.Itoa(diff[3])
			_, ok := cache[key]
			if !ok {
				cache[key] = (next % 10)
			}
			diff = diff[1:]
		}
	}

	for k, v := range cache {
		_, ok := globalCache[k]
		if ok {
			globalCache[k] += v
		} else {
			globalCache[k] = v
		}
	}

	return cur
}

func SolveDay22(input string) string {
	lines := strings.Split(input, "\n")

	result1 := 0
	cache := make(map[string]int)
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err == nil {
			res := calculateSecret(num, cache)
			result1 += res
		}
	}

	result2 := 0
	for _, bananas := range cache {
		if bananas > result2 {
			result2 = bananas
		}
	}
	// 19241711734 2058
	return fmt.Sprintf("%d %d", result1, result2)
}
