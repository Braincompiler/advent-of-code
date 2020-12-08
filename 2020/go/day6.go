package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day6() {
	file, err := os.Open("../inputs/day6.txt")
	if err != nil {
		panic("Failed to open input file for day 6")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var groups [][]string
	for scanner.Scan() {
		line := scanner.Text()
		var persons []string
		for len(line) > 0 {
			persons = append(persons, line)
			scanner.Scan()
			line = scanner.Text()
		}

		groups = append(groups, persons)
	}

	resPuzzle1 := findResultDay6Puzzle1(&groups)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 6: %d", resPuzzle1))

	resPuzzle2 := findResultDay6Puzzle2(&groups)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 6: %d", resPuzzle2))
}

func findResultDay6Puzzle1(groups *[][]string) int {
	sum := 0
	for _, group := range *groups {
		countMap := make(map[int32]bool)
		for _, person := range group {
			for _, q := range person {
				countMap[q] = true
			}
		}
		sum += len(countMap)
	}

	return sum
}

func filterByCount(countMap *map[int32]int) []int32 {
	var newCountMap []int32
	l := len(*countMap)
	for q, n := range *countMap {
		if n == l {
			newCountMap = append(newCountMap, q)
		}
	}

	return newCountMap
}

func findResultDay6Puzzle2(groups *[][]string) int {
	sum := 0
	for _, group := range *groups {
		countMap := make(map[int32]bool)
		for _, person := range group {
			for _, q := range person {
				countMap[q] = true
			}
		}

		var letter []int32
		for k := range countMap {
			letter = append(letter, k)
		}

		for _, person := range group {
			for _, l := range letter {
				if strings.ContainsRune(person, l) {
					sum++
				}
			}
		}
	}

	return sum // 3559
}
