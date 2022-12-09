package main

import (
	"bufio"
	"fmt"
	"os"
)

func day9() {
	file, err := os.Open("../inputs/day9.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 9: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 9: %v", err))
		}
	}(file)

	var input []any

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

	}

	resPuzzle1 := findResultDay9Puzzle1(input)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 9: %d", resPuzzle1))

	resPuzzle2 := findResultDay9Puzzle2(input)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 9: %d", resPuzzle2))
}

func findResultDay9Puzzle1(input []any) int {
	return 0
}

func findResultDay9Puzzle2(input []any) int {
	return 0
}
