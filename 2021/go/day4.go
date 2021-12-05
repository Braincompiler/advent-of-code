package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoNumber struct {
	Number int
	Marked bool
}

type BingoBoard struct {
	Numbers [][]BingoNumber
}

func day4() {
	file, err := os.Open("../inputs/day4.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 4: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 4: %v", err))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	hasNumbers := false
	hasBoards := false
	rowIndex := 0
	currentBoard := CreateBingoBoard()
	var boards []BingoBoard
	var numbers []int
	for scanner.Scan() {
		s := scanner.Text()

		if s == "" {
			if hasBoards {
				boards = append(boards, currentBoard)
				rowIndex = 0
				currentBoard = CreateBingoBoard()
			}
			continue
		}

		if !hasNumbers {
			nums := strings.Split(s, ",")
			numbers = mapSlice(nums, func(s string) int {
				n, _ := strconv.ParseUint(s, 10, 32)
				return int(n)
			})

			hasNumbers = true

			continue
		}

		hasBoards = true

		row := mapSlice(strings.Fields(s), func(s string) int {
			n, _ := strconv.ParseUint(s, 10, 32)
			return int(n)
		})
		currentBoard.SetRow(rowIndex, row)
		rowIndex++
	}

	//fmt.Printf("%v\n%v\n", numbers, boards)

	resPuzzle1 := findResultDay4Puzzle1(numbers, boards)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 4: %d", resPuzzle1))

	resPuzzle2 := findResultDay4Puzzle2(numbers)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 4: %d", resPuzzle2))
}

func findResultDay4Puzzle1(numbers []int, boards []BingoBoard) int {
	weHaveAWinner := false
	var theWinnerBoard BingoBoard
	var theLastNumber int
	for _, n := range numbers {
		for _, board := range boards {
			board.MarkNumber(n)
			if board.CheckWonRowOrColumn() {
				weHaveAWinner = true
				theWinnerBoard = board
				theLastNumber = n
			}
		}

		if weHaveAWinner {
			break
		}
	}

	fmt.Printf("%d = %v", theLastNumber, theWinnerBoard)

	return 0
}

func findResultDay4Puzzle2(numbers []int) int {
	return 0
}

func mapSlice(a []string, f func(string) int) []int {
	var returns []int

	for _, i := range a {
		returns = append(returns, f(i))
	}

	return returns
}

func allSlice(a []BingoNumber, f func(BingoNumber) bool) bool {
	for _, n := range a {
		if !f(n) {
			return false
		}
	}

	return true
}

func CreateBingoNumber(number int) BingoNumber {
	return BingoNumber{
		Number: number,
		Marked: false,
	}
}

func CreateBingoBoard() BingoBoard {
	return BingoBoard{
		Numbers: [][]BingoNumber{
			{CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0)},
			{CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0)},
			{CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0)},
			{CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0)},
			{CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0), CreateBingoNumber(0)},
		},
	}
}

func (bb BingoBoard) SetRow(rowIndex int, row []int) {
	for i, r := range row {
		bb.Numbers[rowIndex][i] = CreateBingoNumber(r)
	}
}

func (bb BingoBoard) MarkNumber(number int) {
	for _, rows := range bb.Numbers {
		for _, cell := range rows {
			if cell.Number == number {
				cell.Marked = true
				continue
			}
		}
	}
}

func (bb BingoBoard) CheckWonRowOrColumn() bool {
	for _, rows := range bb.Numbers {
		hasFullRow := allSlice(rows, func(number BingoNumber) bool {
			return number.Marked
		})
		if hasFullRow {
			return true
		}
	}

	l := len(bb.Numbers)
	colCounter := 0
	for colIndex := 0; colIndex < l; colIndex++ {
		for rowIndex := 0; rowIndex < l; rowIndex++ {
			if bb.Numbers[colIndex][rowIndex].Marked {
				colCounter++
			}
		}
		if colCounter == l {
			return true
		}
	}

	return false
}
