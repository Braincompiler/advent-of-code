package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
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

	input := ""

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		input = s
	}

	resPuzzle1 := findResultDay6Puzzle1(input)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 6: %d", resPuzzle1))

	resPuzzle2 := findResultDay6Puzzle2(input)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 6: %d", resPuzzle2))
}

func findResultDay6Puzzle1(input string) int {
	return findEndpositionOfUniqueBytes(input, 4)
}

func findResultDay6Puzzle2(input string) int {
	return findEndpositionOfUniqueBytes(input, 14)
}

func hasUniqueBytes(r *ring.Ring) bool {
	set := map[int32]bool{}

	r.Do(func(i any) {
		if i != nil {
			set[i.(int32)] = true
		}
	})

	return len(set) == r.Len()
}

func findEndpositionOfUniqueBytes(input string, len int) int {
	result := 0
	r := ring.New(len)

	for i, c := range input {
		r.Value = c
		r = r.Next()

		if hasUniqueBytes(r) {
			result = i
			break
		}
	}

	return result + 1
}
