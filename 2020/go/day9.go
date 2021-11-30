package main

import (
	"bufio"
	"fmt"
	"os"
)

func day9() {
	file, err := os.Open("../inputs/day9.txt")
	if err != nil {
		panic("Failed to open input file for day 9")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int
	for scanner.Scan() {
		line := scanner.Text()

		numbers = append(numbers, parseInt(line))
	}

	resPuzzle1 := findResultDay9Puzzle1(&numbers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 9: %d", resPuzzle1))

	resPuzzle2 := findResultDay9Puzzle2(&numbers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 9: %d", resPuzzle2))
}

func findResultDay9Puzzle1(numbers *[]int) int {
	const PreambleSize = 25
	preamble := make([]int, PreambleSize)
	numSlices := len(*numbers) / PreambleSize

	for i := 0; i < numSlices; i++ {
	}

	return -1
}

func findResultDay9Puzzle2(numbers *[]int) int {
	return -1
}
