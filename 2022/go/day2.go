package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	OpponentRock       uint8 = 'A'
	OpponentPaper      uint8 = 'B'
	OpponentScissor    uint8 = 'C'
	OpponentRockInt          = 1
	OpponentPaperInt         = 2
	OpponentScissorInt       = 3

	MyChoiceRock       uint8 = 'X'
	MyChoicePaper      uint8 = 'Y'
	MyChoiceScissor    uint8 = 'Z'
	MyChoiceRockInt          = 1
	MyChoicePaperInt         = 2
	MyChoiceScissorInt       = 3

	Loose = 0
	Draw  = 3
	Win   = 6

	HaveToLoose uint8 = 'X'
	HaveToDraw  uint8 = 'Y'
	HaveToWin   uint8 = 'Z'
)

type Round struct {
	Opponent uint8
	MyChoice uint8
}

func day2() {
	file, err := os.Open("../inputs/day2.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 1: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 1: %v", err))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rounds []Round
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		rounds = append(rounds, Round{
			Opponent: s[0],
			MyChoice: s[2],
		})
	}

	resPuzzle1 := findResultDay2Puzzle1(rounds)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 2: %d", resPuzzle1))

	resPuzzle2 := findResultDay2Puzzle2(rounds)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 2: %d", resPuzzle2))
}

func findResultDay2Puzzle1(rounds []Round) int {
	sum := 0

	for _, round := range rounds {
		if round.IsMyChoiceWin() {
			sum += round.MyChoiceAsInt() + Win
		} else if round.IsDraw() {
			sum += round.MyChoiceAsInt() + Draw
		} else if round.IsMyChoiceLoose() {
			sum += round.MyChoiceAsInt() + Loose
		}
	}

	return sum
}

func findResultDay2Puzzle2(rounds []Round) int {
	sum := 0

	for _, round := range rounds {
		sum += round.PredictMyChoiceAsInt() + HaveToDoAsInt(round.MyChoice)
	}

	return sum
}

func (r Round) IsDraw() bool {
	return r.OpponentAsInt() == r.MyChoiceAsInt()
}

func (r Round) IsMyChoiceWin() bool {
	return r.IsMyChoiceRock() && r.IsOpponentScissor() ||
		r.IsMyChoicePaper() && r.IsOpponentRock() ||
		r.IsMyChoiceScissor() && r.IsOpponentPaper()
}

func (r Round) IsMyChoiceLoose() bool {
	return !r.IsMyChoiceWin() && !r.IsDraw()
}

func (r Round) OpponentAsInt() int {
	switch r.Opponent {
	case OpponentRock:
		return OpponentRockInt

	case OpponentPaper:
		return OpponentPaperInt

	case OpponentScissor:
		return OpponentScissorInt
	}

	return -1
}

func (r Round) MyChoiceAsInt() int {
	switch r.MyChoice {
	case MyChoiceRock:
		return MyChoiceRockInt

	case MyChoicePaper:
		return MyChoicePaperInt

	case MyChoiceScissor:
		return MyChoiceScissorInt
	}

	return -1
}

func (r Round) IsOpponentRock() bool {
	return r.Opponent == OpponentRock
}

func (r Round) IsOpponentPaper() bool {
	return r.Opponent == OpponentPaper
}

func (r Round) IsOpponentScissor() bool {
	return r.Opponent == OpponentScissor
}

func (r Round) IsMyChoiceRock() bool {
	return r.MyChoice == MyChoiceRock
}

func (r Round) IsMyChoicePaper() bool {
	return r.MyChoice == MyChoicePaper
}

func (r Round) IsMyChoiceScissor() bool {
	return r.MyChoice == MyChoiceScissor
}

func (r Round) PredictMyChoiceAsInt() int {
	if r.MyChoice == HaveToLoose {
		return r.GetMyChoiceToLooseAsInt()
	} else if r.MyChoice == HaveToWin {
		return r.GetMyChoiceToWinAsInt()
	} else if r.MyChoice == HaveToDraw {
		return r.OpponentAsInt()
	}

	return -1
}

func (r Round) GetMyChoiceToLooseAsInt() int {
	if r.Opponent == OpponentRock {
		return MyChoiceScissorInt
	} else if r.Opponent == OpponentPaper {
		return MyChoiceRockInt
	} else if r.Opponent == OpponentScissor {
		return MyChoicePaperInt
	}

	return -1
}

func (r Round) GetMyChoiceToWinAsInt() int {
	if r.Opponent == OpponentRock {
		return MyChoicePaperInt
	} else if r.Opponent == OpponentPaper {
		return MyChoiceScissorInt
	} else if r.Opponent == OpponentScissor {
		return MyChoiceRockInt
	}

	return -1
}

func HaveToDoAsInt(haveToDo uint8) int {
	if haveToDo == HaveToWin {
		return Win
	} else if haveToDo == HaveToDraw {
		return Draw
	} else if haveToDo == HaveToLoose {
		return Loose
	}

	return -1
}
