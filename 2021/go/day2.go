package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	Horizontal int
	Depth      int
}

type PositionWithAim struct {
	Position Position
	Aim      int
}

func day2() {
	file, err := os.Open("../inputs/day2.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 2: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 2: %v", err))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var plannedCourse []Position
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		var waypoint Position

		switch fields[0] {
		case "forward":
			waypoint = Position{
				Horizontal: parseInt(fields[1]),
				Depth:      0,
			}

		case "up":
			waypoint = Position{
				Horizontal: 0,
				Depth:      parseInt(fields[1]) * -1,
			}

		case "down":
			waypoint = Position{
				Horizontal: 0,
				Depth:      parseInt(fields[1]),
			}
		}

		plannedCourse = append(plannedCourse, waypoint)
	}

	resPuzzle1 := findResultDay2Puzzle1(plannedCourse)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 2: %d", resPuzzle1))

	resPuzzle2 := findResultDay2Puzzle2(plannedCourse)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 2: %d", resPuzzle2))
}

func findResultDay2Puzzle1(plannedCourse []Position) int {
	position := Position{
		Horizontal: 0,
		Depth:      0,
	}

	for _, waypoint := range plannedCourse {
		position.Horizontal += waypoint.Horizontal
		position.Depth += waypoint.Depth
	}

	return position.Horizontal * position.Depth
}

func findResultDay2Puzzle2(plannedCourse []Position) int {
	position := PositionWithAim{
		Position: Position{
			Horizontal: 0,
			Depth:      0,
		},
		Aim: 0,
	}

	for _, waypoint := range plannedCourse {
		if waypoint.Horizontal == 0 {
			position.Aim += waypoint.Depth
		}
		position.Position.Horizontal += waypoint.Horizontal
		position.Position.Depth += waypoint.Horizontal * position.Aim
	}

	return position.Position.Horizontal * position.Position.Depth
}
