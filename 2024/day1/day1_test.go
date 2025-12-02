package day1

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestDay1Solution1(t *testing.T) {
	file, err := os.Open("../input/day1/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	day := NewDay1()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		day.Parse(line)
	}

	solution1 := day.Solution1()
	if solution1 != 11 {
		t.Fatalf("Day 1 - Puzzle1: Expected 11 results, got %d", solution1)
	}
}

func TestDay1Solution2(t *testing.T) {
	file, err := os.Open("../input/day1/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	day := NewDay1()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		day.Parse(line)
	}

	solution2 := day.Solution2()
	if solution2 != 31 {
		t.Fatalf("Day 1 - Puzzle2: Expected 31 results, got %d", solution2)
	}
}
