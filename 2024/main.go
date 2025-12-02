package main

import (
	"2024/day1"
	"2024/day2"
	"bufio"
	"fmt"
	"log"
	"os"
)

type AOCDay interface {
	Parse(line string)
	Solution1() int
	Solution2() int
}

const (
	DAYS = 2
)

func main() {
	daySolutions := []AOCDay{
		day1.NewDay1(),
		day2.NewDay2(),
	}

	for day := 1; day <= DAYS; day++ {
		file, err := os.Open(fmt.Sprintf("input/day%d/puzzle1.txt", day))
		if err != nil {
			log.Fatal(err)
		}

		daySolution := daySolutions[day-1]
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			daySolution.Parse(line)
		}

		fmt.Printf("===================== Day #%d =====================\n", day)
		fmt.Printf("Solution 1: %d\n", daySolution.Solution1())
		fmt.Printf("Solution 2: %d\n", daySolution.Solution2())
		fmt.Println("==================================================")

		func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}(file)
	}
}
