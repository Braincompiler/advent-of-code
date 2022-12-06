package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type (
	Instruction struct {
		count int
		from  int
		to    int
	}
)

func day5() {
	file, err := os.Open("../inputs/day5.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 5: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 5: %v", err))
		}
	}(file)

	var instructions []Instruction

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	parseActive := false
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			parseActive = true
			continue
		}

		if parseActive {
			matches := re.FindStringSubmatch(s)

			instructions = append(instructions, Instruction{
				count: parseInt(matches[1]),
				from:  parseInt(matches[2]),
				to:    parseInt(matches[3]),
			})
		}
	}

	resPuzzle1 := findResultDay5Puzzle1(CreateStacks(), instructions)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 5: %s", resPuzzle1))

	resPuzzle2 := findResultDay5Puzzle2(CreateStacks(), instructions)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 5: %s", resPuzzle2))
}

func findResultDay5Puzzle1(stacks []*Stack[uint8], instructions []Instruction) string {
	for _, instruction := range instructions {
		for i := 0; i < instruction.count; i++ {
			value := stacks[instruction.from-1].Pop()
			stacks[instruction.to-1].Push(*value)
		}
	}

	topCrates := ""
	for _, stack := range stacks {
		topCrates += string(*stack.Peek())
	}

	return topCrates
}

func findResultDay5Puzzle2(stacks []*Stack[uint8], instructions []Instruction) string {
	for _, instruction := range instructions {
		values := stacks[instruction.from-1].PopMany(instruction.count)
		reverse(values)
		stacks[instruction.to-1].PushMany(values)
	}

	topCrates := ""
	for _, stack := range stacks {
		topCrates += string(*stack.Peek())
	}

	return topCrates
}

func CreateStacks() []*Stack[uint8] {
	stack1 := NewStack[uint8]().
		Push('L').
		Push('D').
		Push('V').
		Push('T').
		Push('M').
		Push('H').
		Push('F')
	stack2 := NewStack[uint8]().
		Push('H').
		Push('Q').
		Push('G').
		Push('J').
		Push('C').
		Push('T').
		Push('N').
		Push('P')
	stack3 := NewStack[uint8]().
		Push('R').
		Push('S').
		Push('D').
		Push('M').
		Push('P').
		Push('H')
	stack4 := NewStack[uint8]().
		Push('L').
		Push('B').
		Push('V').
		Push('F')
	stack5 := NewStack[uint8]().
		Push('N').
		Push('H').
		Push('G').
		Push('L').
		Push('Q')
	stack6 := NewStack[uint8]().
		Push('W').
		Push('B').
		Push('D').
		Push('G').
		Push('R').
		Push('M').
		Push('P')
	stack7 := NewStack[uint8]().
		Push('G').
		Push('M').
		Push('N').
		Push('R').
		Push('C').
		Push('H').
		Push('L').
		Push('Q')
	stack8 := NewStack[uint8]().
		Push('W').
		Push('L').
		Push('C')
	stack9 := NewStack[uint8]().
		Push('R').
		Push('D').
		Push('L').
		Push('Q').
		Push('J').
		Push('Z').
		Push('M').
		Push('T')

	stacks := []*Stack[uint8]{
		stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9,
	}
	return stacks
}

func CreateSampleStacks() []*Stack[uint8] {
	stack1 := NewStack[uint8]().
		Push('Z').
		Push('N')
	stack2 := NewStack[uint8]().
		Push('M').
		Push('C').
		Push('D')
	stack3 := NewStack[uint8]().
		Push('P')

	stacks := []*Stack[uint8]{
		stack1, stack2, stack3,
	}
	return stacks
}
