package main

import (
	"bufio"
	"fmt"
	"os"
)

func day5() {
	file, err := os.Open("../inputs/day5.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 5: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 5: %v", err))
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

	resPuzzle1 := findResultDay5Puzzle1(input)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 5: %d", resPuzzle1))

	resPuzzle2 := findResultDay5Puzzle2(input)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 5: %d", resPuzzle2))
}

func findResultDay5Puzzle1(input []any) int {
	return 0
}

func findResultDay5Puzzle2(input []any) int {
	return 0
}
