package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func day4() {
	file, err := os.Open("../inputs/day4.txt")
	if err != nil {
		panic("Failed to open input file for day 4")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var passports []map[string]string
	for scanner.Scan() {
		var entries []string
		for {
			line := scanner.Text()
			if len(line) == 0 {
				break
			}

			entries = append(entries, line)
			scanner.Scan()
		}

		passport := make(map[string]string)
		kvPairs := strings.Split(strings.Join(entries, " "), " ")
		for _, kvPair := range kvPairs {
			kv := strings.Split(kvPair, ":")
			passport[kv[0]] = kv[1]
		}

		passports = append(passports, passport)
	}

	file.Close()

	resPuzzle1 := findResultDay4Puzzle1(&passports)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 4: %d", resPuzzle1))

	resPuzzle2 := findResultDay4Puzzle2(&passports)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 4: %d", resPuzzle2))
}

func findResultDay4Puzzle1(passports *[]map[string]string) int {
	numValidPassports := 0
	for _, passport := range *passports {
		lenPassports := len(passport)
		_, cidExists := passport["cid"]
		if lenPassports == 8 || (lenPassports == 7 && !cidExists) {
			numValidPassports++
		}
	}

	return numValidPassports
}

type Validator interface {
	Validate(value string) bool
}

type BetweenValidationRule struct {
	min int
	max int
}

type HeightValidationRule struct {
	minCM int
	maxCM int
	minIN int
	maxIN int
}

type HairColorValidationRule struct {
}

type EyeColorValidationRule struct {
	validColors []string
}

type PassportsIDValidationRule struct {
}

func (r BetweenValidationRule) Validate(value string) bool {
	n, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert %s to an integer number", value))
	}

	return n >= r.min && n <= r.max
}

func (r HeightValidationRule) Validate(value string) bool {
	valueLC := strings.ToLower(value)
	valueNum := valueLC[:len(valueLC)-2]
	n, err := strconv.Atoi(valueNum)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert %s to an integer number", valueNum))
	}

	if strings.HasSuffix(valueLC, "cm") {
		return n >= r.minCM && n <= r.maxCM
	} else if strings.HasSuffix(valueLC, "in") {
		return n >= r.minIN && n <= r.maxIN
	}

	return false
}

func (r EyeColorValidationRule) Validate(value string) bool {
	for _, color := range r.validColors {
		if value == color {
			return true
		}
	}

	return false
}

func (r PassportsIDValidationRule) Validate(value string) bool {
	if len(value) != 9 {
		return false
	}

	for _, c := range value {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

func (r HairColorValidationRule) Validate(value string) bool {
	possibleColorValue := value[1:]
	_, err := strconv.ParseUint(possibleColorValue, 16, 64)

	return value[0] == '#' && len(possibleColorValue) == 6 && err == nil
}

func findResultDay4Puzzle2(passports *[]map[string]string) int {
	validationRules := map[string]Validator{
		"byr": BetweenValidationRule{
			min: 1920,
			max: 2002,
		},
		"iyr": BetweenValidationRule{
			min: 2010,
			max: 2020,
		},
		"eyr": BetweenValidationRule{
			min: 2020,
			max: 2030,
		},
		"hgt": HeightValidationRule{
			minCM: 150,
			maxCM: 193,
			minIN: 59,
			maxIN: 76,
		},
		"hcl": HairColorValidationRule{},
		"ecl": EyeColorValidationRule{validColors: []string{
			"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
		}},
		"pid": PassportsIDValidationRule{},
	}

	numValidPassports := 0
START:
	for _, passport := range *passports {
		lenPassports := len(passport)
		_, cidExists := passport["cid"]
		if lenPassports == 8 || (lenPassports == 7 && !cidExists) {
			for fieldName, validationRule := range validationRules {
				value, hasValue := passport[fieldName]
				if !hasValue || !validationRule.Validate(value) {
					continue START
				}
			}

			numValidPassports++
		}
	}

	return numValidPassports
}
