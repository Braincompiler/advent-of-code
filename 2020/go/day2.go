package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type PasswordEntry struct {
	min      int
	max      int
	letter   rune
	password string
}

func day2() {
	file, err := os.Open("../inputs/day2.txt")
	if err != nil {
		panic("Failed to open input file for day 2")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var passwords []PasswordEntry
	for scanner.Scan() {
		line := scanner.Text()

		passwordEntry := PasswordEntry{
			min:      0,
			max:      0,
			letter:   '0',
			password: "",
		}

		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &passwordEntry.min, &passwordEntry.max, &passwordEntry.letter, &passwordEntry.password)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse line \"%s\"", line))
		}

		passwords = append(passwords, passwordEntry)
	}

	file.Close()

	resPuzzle1 := findResultDay2Puzzle1(&passwords)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 2: %d", resPuzzle1))

	resPuzzle2 := findResultDay2Puzzle2(&passwords)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 2: %d", resPuzzle2))
}

func parsePasswordEntryFieldFunc(c rune) bool {
	return unicode.IsSpace(c) || c == '-' || c == ':'
}

func findResultDay2Puzzle1(passwords *[]PasswordEntry) int {
	numValidPasswords := 0
	for _, passwordEntry := range *passwords {
		letters := strings.FieldsFunc(passwordEntry.password, func(c rune) bool {
			return c != passwordEntry.letter
		})

		lenLetters := len(strings.Join(letters, ""))
		if lenLetters >= passwordEntry.min && lenLetters <= passwordEntry.max {
			numValidPasswords++
		}
	}

	return numValidPasswords
}

func findResultDay2Puzzle2(passwords *[]PasswordEntry) int {
	numValidPasswords := 0
	for _, passwordEntry := range *passwords {
		firstLetterIdx := passwordEntry.min - 1
		secondLetterIdx := passwordEntry.max - 1
		if isValidPasswordPuzzle2(passwordEntry.password, firstLetterIdx, secondLetterIdx, passwordEntry.letter) {
			numValidPasswords++
		}
	}

	return numValidPasswords
}

func isValidPasswordPuzzle2(password string, fl, sl int, letter rune) bool {
	return (rune(password[fl]) == letter) != (rune(password[sl]) == letter)
}
