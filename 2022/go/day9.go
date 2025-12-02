package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type (
	InstructionD9 struct {
		direction uint8
		steps     int
	}
	Coordinate struct {
		x int
		y int
	}
	Rope struct {
		head *Coordinate
		tail *Coordinate

		tailUniqueVisits map[string]bool
	}
)

var (
	DirectionUp    = uint8('U')
	DirectionDown  = uint8('D')
	DirectionLeft  = uint8('L')
	DirectionRight = uint8('R')
)

func day9() {
	file, err := os.Open("../inputs/day9_sample.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open instructions file for day 9: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close instructions file for day 9: %v", err))
		}
	}(file)

	var instructions []InstructionD9

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		parts := strings.Split(s, " ")

		instructions = append(instructions, InstructionD9{
			direction: parts[0][0],
			steps:     parseInt(parts[1]),
		})
	}

	resPuzzle1 := findResultDay9Puzzle1(instructions)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 9: %d", resPuzzle1))

	resPuzzle2 := findResultDay9Puzzle2(instructions)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 9: %d", resPuzzle2))
}

func findResultDay9Puzzle1(instructions []InstructionD9) int {
	rope := Rope{
		head: &Coordinate{
			x: 0,
			y: 0,
		},
		tail: &Coordinate{
			x: 0,
			y: 0,
		},
		tailUniqueVisits: map[string]bool{},
	}

	for _, instr := range instructions {
		rope.MoveHead(instr)
		rope.Dump()
	}

	return len(rope.tailUniqueVisits)
}

func findResultDay9Puzzle2(input []InstructionD9) int {
	return 0
}

func (r Rope) HeadCoversTail() bool {
	return r.head.x == r.tail.x && r.head.y == r.tail.y
}

func (r Rope) MoveHead(i InstructionD9) {
	isOverlapping := r.HeadCoversTail()

	switch i.direction {
	case DirectionUp:
		if r.head.x != r.tail.x {
			r.tail.x = r.head.x
		}
		r.head.y += i.steps
		if isOverlapping {
			r.tail.y += i.steps - 1
		}
		break

	case DirectionDown:
		if r.head.x != r.tail.x {
			r.tail.x = r.head.x
		}
		r.head.y -= i.steps
		if isOverlapping {
			r.tail.y -= i.steps - 1
		}
		break

	case DirectionLeft:
		if r.head.y != r.tail.y {
			r.tail.y = r.head.y
		}
		r.head.x -= i.steps
		if isOverlapping {
			r.tail.x -= i.steps - 1
		}
		break

	case DirectionRight:
		if r.head.y != r.tail.y {
			r.tail.y = r.head.y
		}
		r.head.x += i.steps
		if isOverlapping {
			r.tail.x += i.steps - 1
		}
		break
	}

	r.StoreTailVisit()
}

func (r Rope) Dump() {
	fmt.Printf("H(%d, %d) - T(%d, %d)\n", r.head.x, r.head.y, r.tail.x, r.tail.y)
}

func (r Rope) StoreTailVisit() {
	r.tailUniqueVisits[fmt.Sprintf("(%d,%d)", r.tail.x, r.tail.y)] = true
}
