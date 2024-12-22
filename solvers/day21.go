package solvers

import (
	"aoc-2024/common"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(21, SolveDay21)
}

func calculateNumKeypadMoves(s []int, e []int) string {
	diffY := e[0] - s[0]
	diffX := e[1] - s[1]

	moveX := ""
	if diffX > 0 {
		moveX = moveX + strings.Repeat(">", diffX)
	} else if diffX < 0 {
		moveX = moveX + strings.Repeat("<", -diffX)
	}

	moveY := ""
	if diffY > 0 {
		moveY = moveY + strings.Repeat("v", diffY)
	} else if diffY < 0 {
		moveY = moveY + strings.Repeat("^", -diffY)
	}

	move := ""
	if s[0] == 3 && e[1] == 0 {
		move = moveX + moveY
	} else if s[1] == 0 && e[0] == 3 {
		move = moveY + moveX
	} else if diffX < 0 {
		move = moveY + moveX
	} else {
		move = moveX + moveY
	}
	return move
}

func findKeypadSequence(code []rune, numMap map[rune][]int, start []int) string {
	moves := []string{}
	curX := start[1]
	curY := start[0]
	for i := 0; i < len(code); i++ {
		coordinate := numMap[code[i]]
		move := calculateNumKeypadMoves([]int{curY, curX}, coordinate)
		curX = coordinate[1]
		curY = coordinate[0]
		moves = append(moves, move+"A")
	}
	return strings.Join(moves, "")
}

func calculateDirKeypadMoves(s []int, e []int) string {
	diffY := e[0] - s[0]
	diffX := e[1] - s[1]

	moveX := ""
	if diffX > 0 {
		moveX = moveX + strings.Repeat(">", diffX)
	} else if diffX < 0 {
		moveX = moveX + strings.Repeat("<", -diffX)
	}

	moveY := ""
	if diffY > 0 {
		moveY = moveY + strings.Repeat("v", diffY)
	} else if diffY < 0 {
		moveY = moveY + strings.Repeat("^", -diffY)
	}

	move := ""
	if s[1] == 0 && e[0] == 0 {
		move = moveX + moveY
	} else if s[0] == 0 && e[1] == 0 {
		move = moveY + moveX
	} else if diffX < 0 {
		move = moveX + moveY
	} else {
		move = moveY + moveX
	}
	return move
}

func findDirSequence(sequence []rune, dirMap map[rune][]int, start []int) string {
	moves := []string{}
	curX := start[1]
	curY := start[0]

	for i := range sequence {
		coordinate := dirMap[sequence[i]]
		move := calculateDirKeypadMoves([]int{curY, curX}, coordinate)
		curX = coordinate[1]
		curY = coordinate[0]
		moves = append(moves, move+"A")
	}
	return strings.Join(moves, "")
}

func SolveDay21(input string) string {
	codes := strings.Split(input, "\n")
	numKeypad := map[rune][]int{
		'7': {0, 0},
		'8': {0, 1},
		'9': {0, 2},
		'4': {1, 0},
		'5': {1, 1},
		'6': {1, 2},
		'1': {2, 0},
		'2': {2, 1},
		'3': {2, 2},
		'0': {3, 1},
		'A': {3, 2},
	}
	dirKeypad := map[rune][]int{
		'^': {0, 1},
		'A': {0, 2},
		'<': {1, 0},
		'v': {1, 1},
		'>': {1, 2},
	}

	result1 := 0
	for i := range codes {
		seq1 := findKeypadSequence([]rune(codes[i]), numKeypad, []int{3, 2})
		seq2 := findDirSequence([]rune(seq1), dirKeypad, []int{0, 2})
		seq3 := findDirSequence([]rune(seq2), dirKeypad, []int{0, 2})
		seqLen := len(seq3)
		numerical := strings.Join(strings.Split(codes[i], "A"), "")
		numVal, err := strconv.Atoi(numerical)
		if err == nil {
			fmt.Print(seqLen, "\n")
			result1 += (numVal * seqLen)
		}
	}

	result2 := 0
	return fmt.Sprintf("%d %d", result1, result2)
}
