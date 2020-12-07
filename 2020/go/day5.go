package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func day5() {
	file, err := os.Open("../inputs/day5.txt")
	if err != nil {
		panic("Failed to open input file for day 5")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var seatPartioning []string
	for scanner.Scan() {
		line := scanner.Text()

		seatPartioning = append(seatPartioning, line)
	}

	resPuzzle1 := findResultDay5Puzzle1(&seatPartioning)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 5: %d", resPuzzle1))

	resPuzzle2 := findResultDay5Puzzle2(&seatPartioning)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 5: %d", resPuzzle2))
}

func makeArray(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func findResultDay5Puzzle1(seatPartioning *[]string) int {
	highestId := 0
	for _, sp := range *seatPartioning {
		rows := makeArray(128)
		cols := makeArray(8)
		for i := 0; i < len(sp); i++ {
			c := sp[i]
			halfRowsSize := len(rows) / 2
			if c == 'F' {
				rows = rows[:halfRowsSize]
			} else if c == 'B' {
				rows = rows[halfRowsSize:]
			}

			if i >= 7 {
				halfColsSize := len(cols) / 2
				if c == 'R' {
					cols = cols[:halfColsSize]
				} else if c == 'L' {
					cols = cols[halfColsSize:]
				}
			}

			highestId = max(highestId, rows[0]*8+cols[0])
		}
	}

	return highestId
}

func findResultDay5Puzzle2(seatPartioning *[]string) int {
	seatIds := make([]int, len(*seatPartioning))
	for _, sp := range *seatPartioning {
		seatID := 0
		for i := 0; i <= 9; i++ {
			seatID <<= 1
			var bit int
			switch sp[i] {
			case 'F', 'L':
				bit = 0
			case 'B', 'R':
				bit = 1
			}
			seatID += bit
		}
		seatIds = append(seatIds, seatID)
	}
	//for seatIdx, sp := range *seatPartioning {
	//	rows := makeArray(128)
	//	cols := makeArray(8)
	//	for i := 0; i < len(sp); i++ {
	//		c := sp[i]
	//		halfRowsSize := len(rows) / 2
	//		if c == 'F' {
	//			rows = rows[:halfRowsSize]
	//		} else if c == 'B' {
	//			rows = rows[halfRowsSize:]
	//		}
	//
	//		if i >= 7 {
	//			halfColsSize := len(cols) / 2
	//			if c == 'R' {
	//				cols = cols[:halfColsSize]
	//			} else if c == 'L' {
	//				cols = cols[halfColsSize:]
	//			}
	//		}
	//
	//		seatIds[seatIdx] = rows[0]*8 + cols[0]
	//	}
	//}

	sort.Ints(seatIds)
	//fmt.Printf("%v", seatIds)
	//for i := 0; i < len(seatIds); i++ {
	//	if seatIds[i+1]-seatIds[i] == 2 {
	//		//println(seatIds[i] + 1)
	//		return seatIds[i] + 1
	//	}
	//}

	//left := 0
	//mid := 0
	//right := len(seatIds) - 1
	//for right > left+1 {
	//	mid = (left + right) / 2
	//	if (seatIds[left] - left) != (seatIds[mid] - mid) {
	//		right = mid
	//	} else if (seatIds[right] - right) != (seatIds[mid] - mid) {
	//		left = mid
	//	}
	//}

	missingSeat := 0
	for i, v := range seatIds {
		if i > 0 && seatIds[i-1]+1 != v {
			if v-1 > 0 {
				missingSeat = v - 1
				// fmt.Println("Missing:", v-1)
			}
			//return v - 1
		}
	}

	return missingSeat

	//return seatIds[mid] + 1
}
