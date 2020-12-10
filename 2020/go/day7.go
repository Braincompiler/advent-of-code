package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Suitcase struct {
	color    string
	contains map[string]int
}

const (
	Start    = iota
	Contains = iota
)

func day7() {
	file, err := os.Open("../inputs/day7.txt")
	if err != nil {
		panic("Failed to open input file for day 7")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rxInt := regexp.MustCompile("^(?:[-+]?(?:0|[1-9][0-9]*))$")

	var suitcases []Suitcase
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ',' || r == '.'
		})

		var buffer []string
		var colorName string
		var num int
		state := Start
		contains := make(map[string]int, 3)
		for _, word := range words {
			if strings.HasPrefix(word, "bag") {
				if state == Start {
					colorName = strings.Join(buffer, " ")
					buffer = buffer[:0]
				} else {
					contains[strings.Join(buffer, " ")] = num
					buffer = buffer[:0]
					num = 0
				}

				continue
			}

			if word == "contain" {
				state = Contains
				continue
			}

			if word == "no" {
				state = Start
				contains = nil
				break
			}

			if rxInt.MatchString(word) {
				num, err = strconv.Atoi(word)
				if err != nil {
					panic(fmt.Sprintf("Failed to convert %s to an integer number", word))
				}

				continue
			}

			buffer = append(buffer, word)
		}

		// fmt.Printf("%v", words)

		buffer = buffer[:0]
		suitcases = append(suitcases, Suitcase{
			color:    colorName,
			contains: contains,
		})
	}

	// fmt.Printf("%v", suitcases)

	resPuzzle1 := findResultDay7Puzzle1(&suitcases)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 7: %d", resPuzzle1))

	resPuzzle2 := findResultDay7Puzzle2(&suitcases)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 7: %d", resPuzzle2))
}

func findSuitcases(suitcases *[]Suitcase, innerSuitcases *[]Suitcase) int {
	counter := 0
	otherSuitcases := make(map[string]Suitcase)
	for _, innerSuitcase := range *innerSuitcases {
		for _, suitcase := range *suitcases {
			if suitcase.contains != nil {
				if _, has := suitcase.contains[innerSuitcase.color]; has {
					counter++
					// otherSuitcases = append(otherSuitcases, suitcase)
					otherSuitcases[suitcase.color] = suitcase
				}
			}
		}
	}

	if len(otherSuitcases) == 0 {
		return counter
	}

	var innerInnerSuitcases []Suitcase
	for _, suitcase := range otherSuitcases {
		innerInnerSuitcases = append(innerInnerSuitcases, suitcase)
	}

	return counter + findSuitcases(suitcases, &innerInnerSuitcases)
}

func findResultDay7Puzzle1(suitcases *[]Suitcase) int {
	const mybagColor = "shiny gold"

	counter := 0
	otherSuitcases := make(map[string]Suitcase)
	for _, suitcase := range *suitcases {
		if suitcase.contains != nil {
			if _, has := suitcase.contains[mybagColor]; has {
				counter++
				//otherSuitcases = append(otherSuitcases, suitcase)
				otherSuitcases[suitcase.color] = suitcase
			}
		}
	}

	var innerSuitcases []Suitcase
	for _, suitcase := range otherSuitcases {
		innerSuitcases = append(innerSuitcases, suitcase)
	}

	return counter + findSuitcases(suitcases, &innerSuitcases)
}

func findResultDay7Puzzle2(suitcases *[]Suitcase) int {
	return -1
}
