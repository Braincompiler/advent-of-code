package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var (
	Digit0 = "abcefg"
	Digit1 = "cf"
	Digit2 = "acdeg"
	Digit3 = "acdfg"
	Digit4 = "bcdf"
	Digit5 = "abdfg"
	Digit6 = "abdefg"
	Digit7 = "acf"
	Digit8 = "abcdefg"
	Digit9 = "abcdfg"

	//Digit0Sorted = sortString(Digit0)
	//Digit1Sorted = sortString(Digit1)
	//Digit2Sorted = sortString(Digit2)
	//Digit3Sorted = sortString(Digit3)
	//Digit4Sorted = sortString(Digit4)
	//Digit5Sorted = sortString(Digit5)
	//Digit6Sorted = sortString(Digit6)
	//Digit7Sorted = sortString(Digit7)
	//Digit8Sorted = sortString(Digit8)
	//Digit9Sorted = sortString(Digit9)

	Digit1Len = len(Digit1)
	Digit4Len = len(Digit4)
	Digit7Len = len(Digit7)
	Digit8Len = len(Digit8)

	DigitLens = []int{Digit1Len, Digit4Len, Digit7Len, Digit8Len}
	Digits    = []string{
		Digit0,
		Digit1,
		Digit2,
		Digit3,
		Digit4,
		Digit5,
		Digit6,
		Digit7,
		Digit8,
		Digit9,
	}
	DigitMapping = map[rune]string{
		'0': Digit0,
		'1': Digit1,
		'2': Digit2,
		'3': Digit3,
		'4': Digit4,
		'5': Digit5,
		'6': Digit6,
		'7': Digit7,
		'8': Digit8,
		'9': Digit9,
	}
)

type DisplayLine struct {
	UniqueSignalPatterns    []string
	FourDigitOutputPatterns []string
}

func day8() {
	file, err := os.Open("../inputs/day8_sample.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 8: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 8: %v", err))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []DisplayLine
	for scanner.Scan() {
		s := scanner.Text()

		split := mapStringSlice(strings.Split(s, "|"), func(s string) string {
			return strings.Trim(s, " ")
		})

		lines = append(lines, DisplayLine{
			UniqueSignalPatterns:    strings.Split(split[0], " "),
			FourDigitOutputPatterns: strings.Split(split[1], " "),
		})
	}

	resPuzzle1 := findResultDay8Puzzle1(lines)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 8: %d", resPuzzle1))

	resPuzzle2 := findResultDay8Puzzle2(lines)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 8: %d", resPuzzle2))
}

func findResultDay8Puzzle1(lines []DisplayLine) int {
	counter := 0

	for _, line := range lines {
		for _, pattern := range line.FourDigitOutputPatterns {
			if contains(DigitLens, len(pattern)) {
				counter++
			}
		}
	}

	return counter
}

func findResultDay8Puzzle2(lines []DisplayLine) int {
	// https://github.com/DarioSucic/Advent-of-Code-2021/blob/master/8_2.py

	abc := "abcdefg"
	permutations := permutations([]int{0, 1, 2, 3, 4, 5, 6})

	for _, line := range lines {
		decodingMap := make(map[rune]rune, len(abc))

		for _, permutation := range permutations {
			for _, pair := range zip(abc, permutation) {
				decodingMap[pair.A] = rune(abc[pair.B])
			}
		}

		var decoded bytes.Buffer
		for _, pattern := range line.UniqueSignalPatterns {
			for _, c := range pattern {
				decoded.WriteRune(decodingMap[c])
			}

			if containsString(Digits, sortString(decoded.String())) {
				fmt.Printf("%v\n", decodingMap)
			}
		}
	}

	//fmt.Printf("%v\n", permutations)

	/**
	 0000
	1    2
	1    2
	1    2
	 3333
	4    5
	4    5
	4    5
	 6666
	*/

	//digitConfig := "0000000"
	//digitLen2IndexMap := map[int][]int{
	//	Digit1Len: {2, 5},
	//	Digit4Len: {1, 2, 3, 5},
	//	Digit7Len: {0, 2, 5},
	//	Digit8Len: {0, 1, 2, 3, 4, 5, 6},
	//}
	//
	//for _, line := range lines {
	//	for _, pattern := range line.FourDigitOutputPatterns {
	//		if contains(DigitLens, len(pattern)) {
	//		}
	//	}
	//}

	//otherDigits := []string{Digit0Sorted, Digit2Sorted, Digit3Sorted, Digit5Sorted, Digit6Sorted, Digit9Sorted}
	//counter := 0

	//digits := map[string]string{
	//	Digit0Sorted: "0",
	//	Digit1Sorted: "1",
	//	Digit2Sorted: "2",
	//	Digit3Sorted: "3",
	//	Digit4Sorted: "4",
	//	Digit5Sorted: "5",
	//	Digit6Sorted: "6",
	//	Digit7Sorted: "7",
	//	Digit8Sorted: "8",
	//	Digit9Sorted: "9",
	//}
	//
	//sum := 0
	//for _, line := range lines {
	//	var outputValue bytes.Buffer
	//
	//	for _, pattern := range line.FourDigitOutputPatterns {
	//		outputValue.WriteString(digits[sortString(pattern)])
	//
	//		//if contains(DigitLens, len(pattern)) {
	//		//	counter++
	//		//} else {
	//		//	for _, digit := range otherDigits {
	//		//	}
	//		//}
	//	}
	//
	//	sum += parseInt(outputValue.String())
	//}
	//
	//return sum

	return 0
}
