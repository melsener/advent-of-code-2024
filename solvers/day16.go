package solvers

import (
	"aoc-2024/common"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

func init() {
	common.RegisterSolver(16, SolveDay16)
}

type State struct {
	cost      int
	x, y      int
	direction int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func isValidCell(maze [][]rune, n int, m int, y int, x int) bool {
	return y >= 0 && y < n && x >= 0 && x < m && maze[y][x] != '#'
}

func findLowestCost(maze [][]rune, n int, m int, start []int, end []int) int {
	var dirs = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	cost := make([][][]int, n)
	for i := range cost {
		cost[i] = make([][]int, m)
		for j := range cost[i] {
			cost[i][j] = []int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32}
		}
	}
	path := &PriorityQueue{}
	heap.Init(path)

	initialDirection := 0
	cost[start[0]][start[1]][initialDirection] = 0
	heap.Push(path, &State{cost: 0, x: start[1], y: start[0], direction: initialDirection})

	for path.Len() > 0 {
		current := heap.Pop(path).(*State)
		x, y, currentCost, currentDirection := current.x, current.y, current.cost, current.direction
		if x == end[1] && y == end[0] {
			return currentCost
		}

		for i := range dirs {
			newY := y + dirs[i][0]
			newX := x + dirs[i][1]
			if isValidCell(maze, n, m, newY, newX) {
				moveCost := 1
				turnCost := 0
				if i != currentDirection {
					turnCost = 1000
				}
				newCost := currentCost + moveCost + turnCost
				if newCost < cost[newY][newX][i] {
					cost[newY][newX][i] = newCost
					heap.Push(path, &State{cost: newCost, x: newX, y: newY, direction: i})
				}
			}
		}

	}

	return -1
}

func SolveDay16(input string) string {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	startPos := make([]int, 2)
	endPos := make([]int, 2)
	maze := make([][]rune, n)
	for i := range lines {
		maze[i] = []rune(lines[i])
		s := strings.IndexRune(lines[i], 'S')
		e := strings.IndexRune(lines[i], 'E')
		if s > -1 {
			startPos[0] = i
			startPos[1] = s
		}
		if e > -1 {
			endPos[0] = i
			endPos[1] = e
		}
	}

	result1 := findLowestCost(maze, n, m, startPos, endPos)
	result2 := 0
	// 143580 0
	return fmt.Sprintf("%d %d", result1, result2)
}
