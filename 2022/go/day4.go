package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type SectionIdPair struct {
	from int
	to   int
}

type Section struct {
	first  SectionIdPair
	second SectionIdPair
}

func day4() {
	file, err := os.Open("../inputs/day4.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open sections file for day 4: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close sections file for day 4: %v", err))
		}
	}(file)

	var sections []Section

	re := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		matches := re.FindStringSubmatch(s)

		sections = append(sections, Section{
			first: SectionIdPair{
				from: parseInt(matches[1]),
				to:   parseInt(matches[2]),
			},
			second: SectionIdPair{
				from: parseInt(matches[3]),
				to:   parseInt(matches[4]),
			},
		})
	}

	resPuzzle1 := findResultDay4Puzzle1(sections)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 4: %d", resPuzzle1))

	resPuzzle2 := findResultDay4Puzzle2(sections)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 4: %d", resPuzzle2))
}

func findResultDay4Puzzle1(sections []Section) int {
	total := 0

	for _, section := range sections {
		if section.Contains() {
			total++
		}
	}

	return total
}

func findResultDay4Puzzle2(sections []Section) int {
	total := 0

	for _, section := range sections {
		if section.Overlaps() {
			total++
		}
	}

	return total
}

func (s Section) Contains() bool {
	return s.first.from >= s.second.from && s.first.to <= s.second.to ||
		s.second.from >= s.first.from && s.second.to <= s.first.to
}

func (s Section) Overlaps() bool {
	return IsBetween(s.first.from, s.second.from, s.second.to) ||
		IsBetween(s.first.to, s.second.from, s.second.to) ||
		IsBetween(s.second.from, s.first.from, s.first.to) ||
		IsBetween(s.second.to, s.first.from, s.first.to)
}
