package main

import (
	"bufio"
	"fmt"
	"os"
)

func day8() {
	file, err := os.Open("../inputs/day8.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 8: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 8: %v", err))
		}
	}(file)

	var trees [][]int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		var row []int

		for _, treeHeight := range s {
			row = append(row, parseIntFromChar(treeHeight))
		}

		trees = append(trees, row)
	}

	resPuzzle1 := findResultDay8Puzzle1(trees)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 8: %d", resPuzzle1))

	resPuzzle2 := findResultDay8Puzzle2(trees)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 8: %d", resPuzzle2))
}

func findResultDay8Puzzle1(trees [][]int) int {
	totalVisibleTrees := (len(trees) * 2) + ((len(trees[0]) - 2) * 2)

	for x := 1; x < len(trees)-1; x++ {
		row := trees[x]
		for y := 1; y < len(row)-1; y++ {
			if isTreeVisible(trees, x, y) {
				totalVisibleTrees++
			}
		}
	}

	return totalVisibleTrees
}

func findResultDay8Puzzle2(trees [][]int) int {
	highestScenicScore := 0

	for x := 1; x < len(trees)-1; x++ {
		row := trees[x]
		for y := 1; y < len(row)-1; y++ {
			currentScenicScore := visibleTreesTop(trees, x, y) *
				visibleTreesBottom(trees, x, y) *
				visibleTreesLeft(trees, x, y) *
				visibleTreesRight(trees, x, y)
			highestScenicScore = max(highestScenicScore, currentScenicScore)
		}
	}

	return highestScenicScore
}

func isTreeVisible(trees [][]int, x, y int) bool {
	return visibleFromTop(trees, x, y) ||
		visibleFromBottom(trees, x, y) ||
		visibleFromLeft(trees, x, y) ||
		visibleFromRight(trees, x, y)
}

func visibleFromTop(trees [][]int, x, y int) bool {
	tree := trees[x][y]
	for i := x - 1; i >= 0; i-- {
		if trees[i][y] >= tree {
			return false
		}
	}

	return true
}

func visibleFromBottom(trees [][]int, x, y int) bool {
	tree := trees[x][y]
	for i := x + 1; i < len(trees); i++ {
		if trees[i][y] >= tree {
			return false
		}
	}

	return true
}

func visibleFromLeft(trees [][]int, x, y int) bool {
	tree := trees[x][y]
	for i := y - 1; i >= 0; i-- {
		if trees[x][i] >= tree {
			return false
		}
	}

	return true
}

func visibleFromRight(trees [][]int, x, y int) bool {
	tree := trees[x][y]
	for i := y + 1; i < len(trees[0]); i++ {
		if trees[x][i] >= tree {
			return false
		}
	}

	return true
}

func visibleTreesTop(trees [][]int, x, y int) int {
	tree := trees[x][y]
	total := 0
	for i := x - 1; i >= 0; i-- {
		total++
		if trees[i][y] >= tree {
			return total
		}
	}

	return total
}

func visibleTreesBottom(trees [][]int, x, y int) int {
	tree := trees[x][y]
	total := 0
	for i := x + 1; i < len(trees); i++ {
		total++
		if trees[i][y] >= tree {
			return total
		}
	}

	return total
}

func visibleTreesLeft(trees [][]int, x, y int) int {
	tree := trees[x][y]
	total := 0
	for i := y - 1; i >= 0; i-- {
		total++
		if trees[x][i] >= tree {
			return total
		}
	}

	return total
}

func visibleTreesRight(trees [][]int, x, y int) int {
	tree := trees[x][y]
	total := 0
	for i := y + 1; i < len(trees[0]); i++ {
		total++
		if trees[x][i] >= tree {
			return total
		}
	}

	return total
}
