package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	op      string
	arg     int
	visited bool
}

const (
	NOP = "nop"
	ACC = "acc"
	JMP = "jmp"
)

func day8() {
	file, err := os.Open("../inputs/day8.txt")
	if err != nil {
		panic("Failed to open input file for day 8")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var program []Instruction
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		program = append(program, Instruction{
			op:      parts[0],
			arg:     parseInt(parts[1]),
			visited: false,
		})
	}

	resPuzzle1 := findResultDay8Puzzle1(program)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 8: %d", resPuzzle1)) // 1939

	resPuzzle2 := findResultDay8Puzzle2(program)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 8: %d", resPuzzle2))
}

func runProgram(program []Instruction) (int, bool) {
	acc := 0
	stackPtr := uint(0)
	lastInstIdx := uint(len(program) - 1)

	for {
		inst := &program[stackPtr]
		if inst.visited {
			return acc, false
		}

		switch inst.op {
		case NOP:
			stackPtr++
			break

		case ACC:
			acc += inst.arg
			stackPtr++
			break

		case JMP:
			stackPtr += uint(inst.arg)
			break
		}

		inst.visited = true

		if stackPtr == lastInstIdx {
			break
		}
	}

	return acc, true
}

func resetProgram(program []Instruction) {
	for i := 0; i < len(program); i++ {
		inst := &program[i]
		inst.visited = false
	}
}

func findResultDay8Puzzle1(program []Instruction) int {
	resetProgram(program)
	acc, _ := runProgram(program)

	return acc
}

func findResultDay8Puzzle2(program []Instruction) int {
	for i, inst := range program {
		if inst.op == ACC {
			continue
		}

		resetProgram(program)
		programCopy := make([]Instruction, len(program))
		copy(programCopy, program)
		acc := 0
		successful := false

		switch inst.op {
		case NOP:
			programCopy[i] = Instruction{
				op:  JMP,
				arg: inst.arg,
			}
			break

		case JMP:
			programCopy[i] = Instruction{
				op:  NOP,
				arg: inst.arg,
			}
			break
		}

		acc, successful = runProgram(programCopy)
		if successful {
			return acc
		}
	}

	return -1
}
