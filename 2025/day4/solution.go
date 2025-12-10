package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type day4 struct {
	grid [][]bool
}

func Solve() {
	day := newDay4()
	day.parse()

	fmt.Println("===== Solution for Day 4 =====")
	fmt.Printf("    Part1: %d\n", day.solvePart1())
	fmt.Printf("    Part2: %d\n", day.solvePart2())
}

func newDay4() *day4 {
	return &day4{}
}

func (d *day4) parse() {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitSeq(line, "")

		var row []bool

		for part := range parts {
			if part == "@" {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}

		d.grid = append(d.grid, row)
	}
}

func (d *day4) solvePart1() int {
	sumAccessibleRolls := 0

	for rowIdx, row := range d.grid {
		for cellIdx, cell := range row {
			if cell && numOfAnotherRollsAround(d.grid, rowIdx, cellIdx) < 4 {
				sumAccessibleRolls++
			}
		}
	}

	return sumAccessibleRolls
}

func (d *day4) solvePart2() int {
	totalRemovedRolls := 0

	for {
		sumAccessibleRolls := 0

		for rowIdx, row := range d.grid {
			for cellIdx, cell := range row {
				if cell && numOfAnotherRollsAround(d.grid, rowIdx, cellIdx) < 4 {
					sumAccessibleRolls++
					d.grid[rowIdx][cellIdx] = false
				}
			}
		}

		if sumAccessibleRolls == 0 {
			break
		}

		totalRemovedRolls += sumAccessibleRolls
	}

	return totalRemovedRolls
}

func numOfAnotherRollsAround(grid [][]bool, rowIdx, cellIdx int) int {
	count := 0

	directions := [][2]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // UpLeft
		{-1, 1},  // UpRight
		{1, -1},  // DownLeft
		{1, 1},   // DownRight
	}
	for _, dir := range directions {
		newRowIdx := rowIdx + dir[0]
		newCellIdx := cellIdx + dir[1]
		if isAnotherRoll(grid, newRowIdx, newCellIdx) {
			count++
		}
	}

	return count
}

func isAnotherRoll(grid [][]bool, rowIdx, cellIdx int) bool {
	lenRows := len(grid)
	lenCells := len(grid[0])

	return rowIdx >= 0 && rowIdx < lenRows && cellIdx >= 0 && cellIdx < lenCells && grid[rowIdx][cellIdx]
}
