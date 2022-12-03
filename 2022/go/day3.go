package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rucksack struct {
	Compartment1 string
	Compartment2 string
}

func day3() {
	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 3: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 3: %v", err))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rucksacks []Rucksack
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		halfSize := len(s) / 2
		rucksacks = append(rucksacks, Rucksack{
			Compartment1: s[:halfSize],
			Compartment2: s[halfSize:],
		})
	}

	resPuzzle1 := findResultDay3Puzzle1(rucksacks)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 3: %d", resPuzzle1))

	resPuzzle2 := findResultDay3Puzzle2(rucksacks)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 3: %d", resPuzzle2))
}

func findResultDay3Puzzle1(rucksacks []Rucksack) int32 {
	var sum int32 = 0

	for _, rucksack := range rucksacks {
		sum += rucksack.FindDuplicatesPriority()
	}

	return sum
}

func findResultDay3Puzzle2(rucksacks []Rucksack) int32 {
	var sum int32 = 0

	chunks := chunk(rucksacks, 3)

	for _, rucksackGroup := range chunks {
		sum += findBadgePriority(rucksackGroup)
	}

	return sum
}

func (r *Rucksack) FindDuplicatesPriority() int32 {
	for _, item := range r.Compartment1 {
		if strings.Contains(r.Compartment2, string(item)) {
			return getPriority(item)
		}
	}

	return -1
}

func (r *Rucksack) AllItems() string {
	return r.Compartment1 + r.Compartment2
}

func getPriority(c int32) int32 {
	if c >= 65 && c <= 90 {
		// is upper case letter
		return c - 38
	}

	// can only be lower case letter
	return c - 96
}

func chunk[T any](items []T, size int) (chunks [][]T) {
	for size < len(items) {
		items, chunks = items[size:], append(chunks, items[:size])
	}

	return append(chunks, items)
}

func findBadgePriority(rucksacks []Rucksack) int32 {
	firstRucksackItems := rucksacks[0].AllItems()
	otherRucksacks := rucksacks[1:]

	for _, item := range firstRucksackItems {
		counter := 0
		for _, otherRucksack := range otherRucksacks {
			if strings.Contains(otherRucksack.AllItems(), string(item)) {
				counter++
			}
		}

		if counter == len(otherRucksacks) {
			return getPriority(item)
		}
	}

	return -1
}
