package main

import (
	"bufio"
	"fmt"
	"os"
)

func day3() {
	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		panic("Failed to open input file for day 3")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var treemap [][]int
	for scanner.Scan() {
		line := scanner.Text()
		lineInts := make([]int, len(line))
		for i, c := range line {
			if c == '#' {
				lineInts[i] = 1
			} else {
				lineInts[i] = 0
			}
		}

		treemap = append(treemap, lineInts)
	}

	file.Close()

	resPuzzle1 := findResultDay3Puzzle1(&treemap, 1, 3) // 198
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 3: %d", resPuzzle1))

	resPuzzle2 := findResultDay3Puzzle2(&treemap) // 5140884672
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 3: %d", resPuzzle2))
}

func findResultDay3Puzzle1(treemap *[][]int, incRow, incCol int) int {
	numRows := len(*treemap)
	numCols := len((*treemap)[0])
	row := 0
	col := 0
	numTrees := 0
	for {
		if row >= (numRows - 1) {
			break
		}

		row += incRow
		col = (col + incCol) % numCols

		if (*treemap)[row][col] == 1 {
			numTrees++
		}
	}

	return numTrees
}

func findResultDay3Puzzle2(treemap *[][]int) int {
	return findResultDay3Puzzle1(treemap, 1, 1) *
		findResultDay3Puzzle1(treemap, 1, 3) *
		findResultDay3Puzzle1(treemap, 1, 5) *
		findResultDay3Puzzle1(treemap, 1, 7) *
		findResultDay3Puzzle1(treemap, 2, 1)
}
