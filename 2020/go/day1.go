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
		panic("Failed to open input file for day 1")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbersPuzzle1 []int
	var numbersPuzzle2 []int
	for scanner.Scan() {
		s := scanner.Text()
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %s to an integer number", s))
		}

		numbersPuzzle1 = append(numbersPuzzle1, n)
		if len(s) == 3 {
			numbersPuzzle2 = append(numbersPuzzle2, n)
		}
	}

	file.Close()

	resPuzzle1 := findResultDay1Puzzle1(numbersPuzzle1)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 1: %d", resPuzzle1))

	resPuzzle2 := findResultDay1Puzzle2(numbersPuzzle2)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 1: %d", resPuzzle2))
}

func findResultDay1Puzzle1(numbers []int) int {
	size := len(numbers)
	for i := 0; i < size; i++ {
		prev := numbers[i]
		for j := i + 1; j < size; j++ {
			if prev+numbers[j] == 2020 {
				return prev * numbers[j]
			}
		}
	}

	return -1
}

func findResultDay1Puzzle2(numbers []int) int {
	size := len(numbers)
	for i := 0; i < size; i++ {
		prev := numbers[i]
		for j := i + 1; j < size; j++ {
			prev2 := numbers[j]
			for k := j + 1; k < size; k++ {
				if prev+prev2+numbers[k] == 2020 {
					return prev * prev2 * numbers[k]
				}
			}
		}
	}

	return -1
}
