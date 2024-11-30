package main

import (
	"aoc-2024/common"
	_ "aoc-2024/solvers" // Import the solutions package to trigger registration
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <day>")
	}

	day := os.Args[1]
	dayInt, err := strconv.Atoi(day)
	if err != nil {
		log.Fatalf("Invalid day: %v", day)
	}

	solver, exists := common.GetSolver(dayInt)
	if !exists {
		log.Fatalf("No solver for day %d", dayInt)
	}

	inputFile := filepath.Join("inputs", fmt.Sprintf("day%d.txt", dayInt))
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	result := solver(strings.TrimSpace(string(input)))
	fmt.Printf("Day %d result: %s\n", dayInt, result)
}
