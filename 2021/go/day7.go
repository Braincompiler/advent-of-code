package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day7() {
	file, err := os.Open("../inputs/day7.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 7: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 7: %v", err))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int
	for scanner.Scan() {
		s := scanner.Text()

		numberStrings := strings.Split(s, ",")
		for _, numberString := range numberStrings {
			numbers = append(numbers, parseInt(numberString))
		}
	}

	resPuzzle1 := findResultDay7Puzzle1(numbers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 7: %d", resPuzzle1))

	resPuzzle2 := findResultDay7Puzzle2(numbers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 7: %d", resPuzzle2))
}

func findResultDay7Puzzle1(numbers []int) int {
	return 0
}

func findResultDay7Puzzle2(numbers []int) int {
	return 0
}
