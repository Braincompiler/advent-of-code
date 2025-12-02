package day1

import (
	"braincompiler/aoc2025/utils"
	"bufio"
	"fmt"
	"log"
	"os"
)

type instruction struct {
	dir       string
	numClicks int64
}

type day1 struct {
	instructions []instruction
}

func Solve() {
	day := newDay1()
	day.parse()

	fmt.Println("===== Solution for Day 1 =====")
	fmt.Printf("    Part1: %d\n", day.solvePart1())
	fmt.Printf("    Part2: %d\n", day.solvePart2())
}

func newDay1() *day1 {
	return &day1{}
}

func (d *day1) parse() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		d.instructions = append(d.instructions, instruction{
			dir:       line[:1],
			numClicks: utils.Atoi64(line[1:]),
		})
	}
}

func (d *day1) solvePart1() int64 {
	dial := int64(50)
	num0 := int64(0)

	for _, instr := range d.instructions {
		switch instr.dir {
		case "R":
			dial = (dial + instr.numClicks) % 100
		case "L":
			dial = (dial - instr.numClicks) % 100
		}

		if dial == 0 {
			num0++
		}
	}

	return num0
}

func (d *day1) solvePart2() int {
	dial := 50
	num0 := 0

	for _, instr := range d.instructions {
		for range utils.Range64(0, instr.numClicks) {
			switch instr.dir {
			case "R":
				dial = (dial + 1) % 100
			case "L":
				dial = (dial - 1) % 100
			}

			if dial == 0 {
				num0++
			}
		}
	}

	return num0
}
