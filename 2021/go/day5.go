package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type LineSegment struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rxp, err := regexp.Compile(`(\d+),(\d+)\s+->\s+(\d+),(\d+)`)
	if err != nil {
		panic(fmt.Sprintf("Failed to compile regexp: %v", err))
	}

	var lineSegments []LineSegment
	for scanner.Scan() {
		s := scanner.Text()

		matches := rxp.FindStringSubmatch(s)

		lineSegments = append(lineSegments, LineSegment{
			x1: parseInt(matches[1]),
			y1: parseInt(matches[2]),
			x2: parseInt(matches[3]),
			y2: parseInt(matches[4]),
		})
	}

	resPuzzle1 := findResultDay5Puzzle1(lineSegments)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 5: %d", resPuzzle1))

	resPuzzle2 := findResultDay5Puzzle2(lineSegments)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 5: %d", resPuzzle2))
}

func findResultDay5Puzzle1(lineSegments []LineSegment) int {
	diagram := createDiagram(lineSegments)
	segments := filterHorizontalAndVerticalLines(lineSegments)

	for _, segment := range segments {
		diagram = segment.MarkLine(diagram)
	}

	//dumpDiagram(diagram)

	return countOverlaps(diagram)
}

func findResultDay5Puzzle2(lineSegments []LineSegment) int {
	diagram := createDiagram(lineSegments)
	segments := filterHorizontalVerticalAndDiagonalLines(lineSegments)

	for _, segment := range segments {
		diagram = segment.MarkLine(diagram)
	}

	//dumpDiagram(diagram)

	return countOverlaps(diagram)
}

func filterHorizontalAndVerticalLines(lineSegments []LineSegment) []LineSegment {
	var result []LineSegment

	for _, segment := range lineSegments {
		if segment.x1 == segment.x2 || segment.y1 == segment.y2 {
			result = append(result, segment)
		}
	}

	return result
}

func filterHorizontalVerticalAndDiagonalLines(lineSegments []LineSegment) []LineSegment {
	var result []LineSegment

	for _, segment := range lineSegments {
		if segment.x1 == segment.x2 || segment.y1 == segment.y2 || abs(segment.x1-segment.x2) == abs(segment.y1-segment.y2) {
			result = append(result, segment)
		}
	}

	return result
}

func getMaxX(lineSegments []LineSegment) int {
	result := 0

	for _, segment := range lineSegments {
		result = max(result, segment.x1, segment.x2)
	}

	return result
}

func getMaxY(lineSegments []LineSegment) int {
	result := 0

	for _, segment := range lineSegments {
		result = max(result, segment.y1, segment.y2)
	}

	return result
}

func createDiagram(lineSegments []LineSegment) [][]int {
	maxXLen := getMaxX(lineSegments) + 1
	maxYLen := getMaxY(lineSegments) + 1

	diagram := make([][]int, maxXLen)
	for i := 0; i < maxXLen; i++ {
		diagram[i] = make([]int, maxYLen)
	}

	return diagram
}

func dumpDiagram(diagram [][]int) {
	var sb strings.Builder

	for _ /*columnIndex*/, rows := range diagram {
		for _ /*rowIndex*/, cell := range rows {
			//sb.WriteString(fmt.Sprintf("r=%d,c=%d", rowIndex, columnIndex))
			sb.WriteString(formatCell(cell))
		}

		sb.WriteString("\n")
	}

	println(sb.String())
}

func formatCell(cell int) string {
	if cell == 0 {
		return "."
	}

	return strconv.Itoa(cell)
}

func getFromTo(a, b int) (int, int) {
	if a > b {
		return b, a
	}

	return a, b
}

//func (ls LineSegment) MarkLine(diagram [][]int) [][]int {
//	xa, xb := getFromTo(ls.x1, ls.x2)
//	ya, yb := getFromTo(ls.y1, ls.y2)
//
//	if xa == xb {
//		for i := ya; i <= yb; i++ {
//			diagram[i][xa]++
//		}
//	} else {
//		for i := xa; i <= xb; i++ {
//			diagram[ya][i]++
//		}
//	}
//
//	return diagram
//}

func (ls LineSegment) MarkLine(diagram [][]int) [][]int {
	xa, xb := getFromTo(ls.x1, ls.x2)
	ya, yb := getFromTo(ls.y1, ls.y2)

	if xa == xb {
		for i := ya; i <= yb; i++ {
			diagram[i][xa]++
		}
	} else if ya == yb {
		for i := xa; i <= xb; i++ {
			diagram[ya][i]++
		}
	} else {
		//   0 1 2 3 4 5 6 7 8 9
		// 0 o . . . . . . . . .
		// 1 . o . . . . . . . .
		// 2 x . o . . . . . . .
		// 3 . x . o . . . . . .
		// 4 . . x . o . . . . .
		// 5 . . . x . o . . . .
		// 6 . . . . x . o . . .
		// 7 . . . . . . . o . .
		// 8 . . . . . . . . o .
		// 9 . . . . . . . . . .

		xa = ls.x1
		ya = ls.y1
		xb = ls.x2
		yb = ls.y2
		diffX := ls.x2 - ls.x1
		diffY := ls.y2 - ls.y1
		stepX := 1
		stepY := 1
		if diffX < 0 {
			stepX *= -1
		}
		if diffY < 0 {
			stepY *= -1
		}

		for {
			diagram[ya][xa]++

			if xa == xb || ya == yb {
				break
			}

			xa += stepX
			ya += stepY
		}
	}

	return diagram
}

func countOverlaps(diagram [][]int) int {
	count := 0

	for _, rows := range diagram {
		for _, cell := range rows {
			if cell > 1 {
				count++
			}
		}
	}

	return count
}
