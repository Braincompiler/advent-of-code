package main

import (
	"bufio"
	"fmt"
	"math"
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
	minSumOfFuel := math.MaxInt
	for i := 1; i <= len(numbers); i++ {
		sumPerHorizontalPosition := 0
		for _, crab := range numbers {
			sumPerHorizontalPosition += abs(i - crab)
		}

		//fmt.Printf("%d: %d\n", i, sumPerHorizontalPosition)
		minSumOfFuel = min(minSumOfFuel, sumPerHorizontalPosition)
	}

	return minSumOfFuel
}

func findResultDay7Puzzle2(numbers []int) int {
	minSumOfFuel := math.MaxInt
	for i := 1; i <= len(numbers); i++ {
		sumOfStep := 0
		for _, crab := range numbers {
			sumOfStep += sumSteps(crab, i)
		}

		minSumOfFuel = min(minSumOfFuel, sumOfStep)
	}

	return minSumOfFuel
}

func sumSteps(crab, steps int) int {
	sum := 0

	for i := 1; i <= abs(crab-steps); i++ {
		sum += i
	}

	return sum
}
