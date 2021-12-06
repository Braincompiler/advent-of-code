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

	boards = append(boards, currentBoard)

	//fmt.Printf("%v\n%v\n", numbers, boards)

	resetBoards(boards)
	resPuzzle1 := findResultDay4Puzzle1(numbers, boards)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 4: %d", resPuzzle1))

	resetBoards(boards)
	resPuzzle2 := findResultDay4Puzzle2(numbers, boards)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 4: %d", resPuzzle2))
}

func findResultDay4Puzzle1(numbers []int, boards []BingoBoard) int {
	weHaveAWinner := false
	var theWinnerBoard BingoBoard
	var theLastNumber int
	for _, n := range numbers {
		for _, board := range boards {
			board.MarkNumber(n)
			if board.HasWonRowOrColumn() {
				weHaveAWinner = true
				theWinnerBoard = board
				theLastNumber = n
			}
		}

		if weHaveAWinner {
			break
		}
	}

	fmt.Printf("%d = %d = %d\n", theLastNumber, sum(theWinnerBoard.GetUnmarkedNumbers()), theLastNumber*sum(theWinnerBoard.GetUnmarkedNumbers()))

	return theLastNumber * sum(theWinnerBoard.GetUnmarkedNumbers())
}

func findResultDay4Puzzle2(numbers []int, boards []BingoBoard) int {
	var theWinnerBoard BingoBoard
	var theLastNumber int
	weHaveAWinner := false

	for ni, n := range numbers {
		println(ni)
		for i, board := range boards {
			board.MarkNumber(n)
			if board.HasWonRowOrColumn() {
				boards = append(boards[:i], boards[i+1:]...)

				// Debug to see the problem and fix it

				if len(boards) == 0 {
					theWinnerBoard = board
					theLastNumber = n
					weHaveAWinner = true
				}

				break
			}
		}

		if weHaveAWinner {
			break
		}
	}

	fmt.Printf("%d = %d = %d\n", theLastNumber, sum(theWinnerBoard.GetUnmarkedNumbers()), theLastNumber*sum(theWinnerBoard.GetUnmarkedNumbers()))

	return theLastNumber * sum(theWinnerBoard.GetUnmarkedNumbers())
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

func sum(a []int) int {
	result := 0
	for _, number := range a {
		result += number
	}

	return result
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

func resetBoards(boards []BingoBoard) {
	for _, board := range boards {
		board.Reset()
	}
}

func (bb *BingoBoard) SetRow(rowIndex int, row []int) {
	for i, r := range row {
		bb.Numbers[rowIndex][i] = CreateBingoNumber(r)
	}
}

func (bb *BingoBoard) MarkNumber(number int) {
	for rowIndex, row := range bb.Numbers {
		for cellIndex, cell := range row {
			if cell.Number == number {
				bb.Numbers[rowIndex][cellIndex].Marked = true
				continue
			}
		}
	}
}

func (bb *BingoBoard) HasWonRowOrColumn() bool {
	//fmt.Printf("%v\n", bb)
	for _, row := range bb.Numbers {
		hasFullRow := allSlice(row, func(number BingoNumber) bool {
			return number.Marked
		})
		if hasFullRow {
			return true
		}
	}

	l := len(bb.Numbers)
	counter := 0
	for rowIndex := 0; rowIndex < l; rowIndex++ {
		for colIndex := 0; colIndex < l; colIndex++ {
			bingoNumber := bb.Numbers[colIndex][rowIndex]
			if bingoNumber.Marked {
				counter++
			}
		}
		if counter == l {
			return true
		}

		counter = 0
	}

	return false
}

func (bb *BingoBoard) GetUnmarkedNumbers() []int {
	var unmarkedNumbers []int

	for _, row := range bb.Numbers {
		for _, cell := range row {
			if !cell.Marked {
				unmarkedNumbers = append(unmarkedNumbers, cell.Number)
			}
		}
	}

	return unmarkedNumbers
}

func (bb *BingoBoard) Reset() {
	for rowIndex, row := range bb.Numbers {
		for cellIndex := range row {
			bb.Numbers[rowIndex][cellIndex].Marked = false
		}
	}
}
