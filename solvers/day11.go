package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(11, SolveDay11)
}

func blink(nums []int) []int {
	res := []int{}
	for i := range nums {
		if nums[i] == 0 {
			res = append(res, 1)
			continue
		}

		str := strconv.Itoa(nums[i])
		if len(str)%2 == 0 {
			left := str[:len(str)/2]
			right := str[len(str)/2:]

			lnum, err1 := strconv.Atoi(left)
			rnum, err2 := strconv.Atoi(right)

			if err1 == nil && err2 == nil {
				res = append(res, lnum)
				res = append(res, rnum)
			}
		} else {
			res = append(res, nums[i]*2024)
		}
	}
	return res
}

func blink2(num int, turn int, cache map[string]int) int {
	key := strconv.Itoa(turn) + ":" + strconv.Itoa(num)
	value, exists := cache[key]
	if exists {
		return value
	}

	if turn == 0 {
		return 1
	}

	str := strconv.Itoa(num)

	result := 0
	if num == 0 {
		result += blink2(1, turn-1, cache)
	} else if len(str)%2 == 0 {
		left := str[:len(str)/2]
		right := str[len(str)/2:]

		lnum, err1 := strconv.Atoi(left)
		rnum, err2 := strconv.Atoi(right)

		if err1 == nil && err2 == nil {
			result += blink2(lnum, turn-1, cache)
			result += blink2(rnum, turn-1, cache)
		}
	} else {
		result += blink2(num*2024, turn-1, cache)
	}

	cache[key] = result
	return result
}

func SolveDay11(input string) string {
	nn := strings.Split(input, " ")
	nums := []int{}
	nums2 := []int{}
	for j := range nn {
		num, err := strconv.Atoi(nn[j])
		if err == nil {
			nums = append(nums, num)
			nums2 = append(nums2, num)
		}
	}

	fmt.Print(nums, "\n")
	for i := 0; i < 25; i++ {
		nums = blink(nums)
	}
	result1 := len(nums)
	// Part 2
	result2 := 0

	cache := make(map[string]int)
	for i := range nums2 {
		result2 += blink2(nums2[i], 75, cache)
	}
	return fmt.Sprintf("%d %d", result1, result2)
}
