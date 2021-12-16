package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	Digit0 = "deagbc"
	Digit1 = "eg"
	Digit2 = "dafgc"
	Digit3 = "dafbc"
	Digit4 = "eafb"
	Digit5 = "defbc"
	Digit6 = "defgbc"
	Digit7 = "dab"
	Digit8 = "deafgbc"
	Digit9 = "deafbc"

	Digit0Sorted = sortString(Digit0)
	Digit1Sorted = sortString(Digit1)
	Digit2Sorted = sortString(Digit2)
	Digit3Sorted = sortString(Digit3)
	Digit4Sorted = sortString(Digit4)
	Digit5Sorted = sortString(Digit5)
	Digit6Sorted = sortString(Digit6)
	Digit7Sorted = sortString(Digit7)
	Digit8Sorted = sortString(Digit8)
	Digit9Sorted = sortString(Digit9)

	Digit1Len = len(Digit1)
	Digit4Len = len(Digit4)
	Digit7Len = len(Digit7)
	Digit8Len = len(Digit8)

	DigitLens = []int{Digit1Len, Digit4Len, Digit7Len, Digit8Len}
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
	// 1, 4, 7, 8 => use these digits to find the display configuration (from UniqueSignalPatterns) and then map the digits from DigitOutputPatterns

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
