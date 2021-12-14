package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	InitTimerValueNewLanternfish  = 8
	ResetTimerValueOldLanternfish = 6
	DaysTotalPuzzle1              = 80
	DaysTotalPuzzle2              = 256
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

	var lanternFishTimers []int
	for scanner.Scan() {
		s := scanner.Text()

		numberStrings := strings.Split(s, ",")
		for _, numberString := range numberStrings {
			lanternFishTimers = append(lanternFishTimers, parseInt(numberString))
		}
	}

	resPuzzle1 := findResultDay6Puzzle1(lanternFishTimers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 6: %d", resPuzzle1))

	resPuzzle2 := findResultDay6Puzzle2(lanternFishTimers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 6: %d", resPuzzle2))
}

func findResultDay6Puzzle1(lanternFishTimers []int) int {
	var lanternFishTimersP1 []int
	for _, lanternFishTimer := range lanternFishTimers {
		lanternFishTimersP1 = append(lanternFishTimersP1, lanternFishTimer)
	}

	for day := 0; day < DaysTotalPuzzle1; day++ {
		lanternFishTimersP1 = decreaseTimersAndCreateNewLanternFishes(lanternFishTimersP1)
	}

	return len(lanternFishTimersP1)
}

func findResultDay6Puzzle2(lanternFishTimers []int) uint64 {
	counter := createCounter(lanternFishTimers)

	for day := 0; day < DaysTotalPuzzle2; day++ {
		numZeros := counter[0]

		for i := 1; i <= 10; i++ {
			counter[i-1] = counter[i]
		}

		counter[ResetTimerValueOldLanternfish] += numZeros
		counter[InitTimerValueNewLanternfish] += numZeros
	}

	return sumCounter(counter)
}

func decreaseTimersAndCreateNewLanternFishes(lanternFishTimers []int) []int {
	newLanternFishCounter := 0

	for i := 0; i < len(lanternFishTimers); i++ {
		if lanternFishTimers[i] == 0 {
			lanternFishTimers[i] = ResetTimerValueOldLanternfish
			newLanternFishCounter++
		} else {
			lanternFishTimers[i]--
		}
	}

	for i := 0; i < newLanternFishCounter; i++ {
		lanternFishTimers = append(lanternFishTimers, InitTimerValueNewLanternfish)
	}

	return lanternFishTimers
}

func createCounter(numbers []int) map[int]int {
	counter := make(map[int]int)

	for _, number := range numbers {
		counter[number]++
	}

	return counter
}

func sumCounter(counter map[int]int) uint64 {
	var sum uint64

	for _, count := range counter {
		sum += uint64(count)
	}

	return sum
}
