package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type bank struct {
	batteries []int64
}

type day3 struct {
	banks []bank
}

func Solve() {
	day := newDay3()
	day.parse()

	fmt.Println("===== Solution for Day 3 =====")
	fmt.Printf("    Part1: %d\n", day.solvePart1())
	fmt.Printf("    Part2: %d\n", day.solvePart2())
}

func newDay3() *day3 {
	return &day3{}
}

func (d *day3) parse() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bank := bank{}

		for _, joltage := range line {
			bank.batteries = append(bank.batteries, int64(joltage-'0'))
		}

		d.banks = append(d.banks, bank)
	}
}

func (d *day3) solvePart1() int64 {
	finalJoltage := int64(0)

	for _, bank := range d.banks {
		bankMaxJoltage := int64(0)
		for i, battJoltage := range bank.batteries {
			for j := i + 1; j < len(bank.batteries); j++ {
				bankMaxJoltage = max(bankMaxJoltage, int64(battJoltage*10+bank.batteries[j]))
			}
		}

		finalJoltage += bankMaxJoltage
	}

	return finalJoltage
}

func (d *day3) solvePart2() int64 {
	finalJoltage := int64(0)

	for _, bank := range d.banks {
		bankMaxJoltage := int64(0)
		dm := int64(-1)
		for i := int64(1); i <= 12; i++ {
			dm++
			for j := dm + 1; j < int64(len(bank.batteries))-12+i; j++ {
				if bank.batteries[j] > bank.batteries[dm] {
					dm = j
				}
			}

			bankMaxJoltage = bankMaxJoltage*10 + bank.batteries[dm]
		}

		finalJoltage += bankMaxJoltage
	}

	return finalJoltage
}
