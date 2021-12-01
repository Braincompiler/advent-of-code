package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day1() {
	file, err := os.Open("../inputs/day1.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 1: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 1: %v", err))
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

	resPuzzle1 := findResultDay1Puzzle1(numbers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 1: %d", resPuzzle1))

	resPuzzle2 := findResultDay1Puzzle2(numbers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 1: %d", resPuzzle2))
}

func findResultDay1Puzzle1(depths []int) int {
	increases := 0
	beforeDepth := 0

	for i, depth := range depths {
		if i > 0 && depth > beforeDepth {
			increases++
		}

		beforeDepth = depth
	}

	return increases
}

func findResultDay1Puzzle2(numbers []int) int {
	var chunks [][]int

	for i := 0; i < len(numbers); i++ {
		if len(numbers[i:]) >= 3 {
			chunks = append(chunks, numbers[i:i+3])
		}
	}

	var sums []int
	for _, chunk := range chunks {
		sums = append(sums, chunk[0]+chunk[1]+chunk[2])
	}

	return findResultDay1Puzzle1(sums)
}
