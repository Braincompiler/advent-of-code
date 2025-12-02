package day2

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestDay2Solution1(t *testing.T) {
	file, err := os.Open("../input/day2/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	day := NewDay2()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		day.Parse(line)
	}

	numReports := len(day.reports)
	if numReports != 6 {
		t.Fatalf("Day 2 - Puzzle1: Expected 6 reports, got %d", numReports)
	}

	for _, report := range day.reports {
		numLevels := len(report.levels)
		if numLevels != 5 {
			t.Fatalf("Day 2 - Puzzle1: Expected 5 levels, got %d", numLevels)
		}
	}

	solution1 := day.Solution1()
	if solution1 != 2 {
		t.Fatalf("Day 2 - Puzzle1: Expected 2 reports are safe, got %d", solution1)
	}
}

func TestDay2Solution2(t *testing.T) {
	file, err := os.Open("../input/day2/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	day := NewDay2()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		day.Parse(line)
	}

	solution2 := day.Solution2()
	if solution2 != 4 {
		t.Fatalf("Day 2 - Puzzle2: Expected 4 results, got %d", solution2)
	}
}
