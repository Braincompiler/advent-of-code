package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day6() {
	file, err := os.Open("../inputs/day6.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 6: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 6: %v", err))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int
	for scanner.Scan() {
		s := scanner.Text()
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %s to an integer number: %v", s, err))
		}

		numbers = append(numbers, n)
	}

	resPuzzle1 := findResultDay6Puzzle1(numbers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 6: %d", resPuzzle1))

	resPuzzle2 := findResultDay6Puzzle2(numbers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 6: %d", resPuzzle2))
}

func findResultDay6Puzzle1(numbers []int) int {
	return 0
}

func findResultDay6Puzzle2(numbers []int) int {
	return 0
}
