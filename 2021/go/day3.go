package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func day3() {
	file, err := os.Open("../inputs/day3.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 3: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 3: %v", err))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var diagnosticBinaries [][]int
	for scanner.Scan() {
		s := scanner.Text()

		var binary []int
		for _, c := range s {
			binary = append(binary, charBitToIntBit(c))
		}

		diagnosticBinaries = append(diagnosticBinaries, binary)
	}

	resPuzzle1 := findResultDay3Puzzle1(diagnosticBinaries) // 2498354
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 3: %d", resPuzzle1))

	resPuzzle2 := findResultDay3Puzzle2(diagnosticBinaries)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 3: %d", resPuzzle2))
}

func findResultDay3Puzzle1(diagnosticBinaries [][]int) int64 {
	counters := countTheBits(diagnosticBinaries)

	//fmt.Printf("Counters: %v\n", counters)

	var gammaBinary bytes.Buffer
	var epsilonBinary bytes.Buffer
	for _, counter := range counters {
		bit := getMostCommonBit(counter)
		gammaBinary.WriteRune(bit)
		epsilonBinary.WriteRune(getCounterBit(bit))
	}

	gamma := binaryToDecimalString(gammaBinary)
	epsilon := binaryToDecimalString(epsilonBinary)

	//fmt.Printf("Gamma: %s = %d\n", gammaBinary.String(), gamma)
	//fmt.Printf("Epsilon: %s = %d\n", epsilonBinary.String(), epsilon)

	return gamma * epsilon
}

func findResultDay3Puzzle2(diagnosticBinaries [][]int) int64 {
	oxygenGeneratorRatingSlice := findBinaryByBitCriteria(diagnosticBinaries, 1, func(counter []int) int {
		if counter[1] > counter[0] || counter[1] == counter[0] {
			return 1
		}

		return 0
	})
	co2ScrubberRatingSlice := findBinaryByBitCriteria(diagnosticBinaries, 0, func(counter []int) int {
		if counter[0] < counter[1] || counter[1] == counter[0] {
			return 0
		}

		return 1
	})

	//fmt.Printf("oxygenGeneratorRatingSlice = %v\n", oxygenGeneratorRatingSlice)
	//fmt.Printf("co2ScrubberRatingSlice = %v\n", co2ScrubberRatingSlice)

	var oxygenGeneratorRatingBinary bytes.Buffer
	for _, i := range oxygenGeneratorRatingSlice {
		oxygenGeneratorRatingBinary.WriteRune(intBitToCharBit(i))
	}

	var co2ScrubberRatingBinary bytes.Buffer
	for _, i := range co2ScrubberRatingSlice {
		co2ScrubberRatingBinary.WriteRune(intBitToCharBit(i))
	}

	oxygenGeneratorRating := binaryToDecimalString(oxygenGeneratorRatingBinary)
	co2ScrubberRating := binaryToDecimalString(co2ScrubberRatingBinary)

	//fmt.Printf("oxygenGeneratorRating = %d\n", oxygenGeneratorRating)
	//fmt.Printf("co2ScrubberRating = %d\n", co2ScrubberRating)

	return oxygenGeneratorRating * co2ScrubberRating
}

type MatchCounter func([]int) int

func findBinaryByBitCriteria(diagnosticBinaries [][]int, criteria int, f MatchCounter) []int {
	bitIndex := 0
	maxBitIndex := len(diagnosticBinaries[0]) - 1
	for true {
		counters := countTheBits(diagnosticBinaries)
		n := f(counters[bitIndex])
		diagnosticBinaries = filter(diagnosticBinaries, func(ints []int) bool {
			return ints[bitIndex] == n
		})

		if len(diagnosticBinaries) == 1 {
			break
		}

		if bitIndex < maxBitIndex {
			bitIndex++
		} else {
			bitIndex = 0
		}
	}

	//fmt.Printf("diagnosticBinaries (%d) = %v\n", criteria, diagnosticBinaries)

	return diagnosticBinaries[0]
}

func countTheBits(diagnosticBinaries [][]int) (counters [][]int) {
	counters = [][]int{}

	for m, binary := range diagnosticBinaries {
		for i, bit := range binary {
			if m == 0 {
				counters = append(counters, []int{0, 0})
			}

			counters[i][bit]++
		}
	}

	return
}

func getMostCommonBit(counter []int) rune {
	if counter[1] > counter[0] {
		return '1'
	}

	return '0'
}

func getCounterBit(bit rune) rune {
	if bit == '1' {
		return '0'
	}

	return '1'
}

func charBitToIntBit(c int32) int {
	if c == '1' {
		return 1
	}

	return 0
}

func intBitToCharBit(n int) rune {
	if n == 1 {
		return '1'
	}

	return '0'
}

type FilterFunc func([]int) bool

func filter(numbers [][]int, pred FilterFunc) [][]int {
	var result [][]int

	for _, n := range numbers {
		if pred(n) {
			result = append(result, n)
		}
	}

	return result
}

func binaryToDecimalString(binary bytes.Buffer) int64 {
	n, err := strconv.ParseInt(binary.String(), 2, 0)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse binary %v to int: %v", binary, err))
	}

	return n
}
