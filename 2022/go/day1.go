package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	calories []int
}

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

	var elves []Elf
	var calories []int
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			elves = append(elves, Elf{
				calories: calories,
			})

			calories = nil

			continue
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %s to an integer number: %v", s, err))
		}

		calories = append(calories, n)
	}

	resPuzzle1 := findResultDay1Puzzle1(elves)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 1: %d", resPuzzle1))

	resPuzzle2 := findResultDay1Puzzle2(elves)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 1: %d", resPuzzle2))
}

func findResultDay1Puzzle1(elves []Elf) int {
	maxTotalCalories := 0

	for _, elf := range elves {
		maxTotalCalories = Max(maxTotalCalories, elf.TotalCalories())
	}

	return maxTotalCalories
}

func findResultDay1Puzzle2(elves []Elf) int {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCalories() > elves[j].TotalCalories()
	})

	return elves[0].TotalCalories() + elves[1].TotalCalories() + elves[2].TotalCalories()
}

func (e Elf) TotalCalories() int {
	sum := 0

	for _, calory := range e.calories {
		sum += calory
	}

	return sum
}

func Max(a, b int) int {
	if a == b || a > b {
		return a
	}

	return b
}
