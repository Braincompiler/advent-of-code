package day6

import (
	"braincompiler/aoc2025/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type column struct {
	values   []int64
	operator string
}

type day6 struct {
	columns  []column
	columns2 []column
}

func Solve() {
	day := newDay6()
	day.parse()

	fmt.Println("===== Solution for Day 6 =====")
	fmt.Printf("    Part1: %d\n", day.solvePart1())
	fmt.Printf("    Part2: %d\n", day.solvePart2())
}

func newDay6() *day6 {
	return &day6{}
}

func (d *day6) parse() {
	file, err := os.Open("./day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		parts := strings.Fields(line)

		if len(d.columns) == 0 {
			for i := 0; i < len(parts); i++ {
				d.columns = append(d.columns, column{})
			}
		}

		firstChar := parts[0]
		if firstChar == "*" || firstChar == "+" {
			for colIdx, op := range parts {
				d.columns[colIdx].operator = op
			}
		} else {
			for colIdx, num := range parts {
				n := utils.Atoi64(strings.Trim(num, " \t"))
				d.columns[colIdx].values = append(d.columns[colIdx].values, n)
			}
		}
	}

	builder := strings.Builder{}
	col := column{}
	for cIdx := range lines[0] {
		for _, line := range lines {
			if line[cIdx] == '+' || line[cIdx] == '*' {
				col.operator = string(line[cIdx])
				continue
			}

			if line[cIdx] != ' ' {
				builder.WriteByte(line[cIdx])
			}
		}

		if builder.Len() > 0 {
			col.values = append(col.values, utils.Atoi64(builder.String()))
			builder.Reset()
		} else {
			d.columns2 = append(d.columns2, col)
			col = column{}
		}
	}
	d.columns2 = append(d.columns2, col)
}

func (d *day6) solvePart1() int64 {
	return d.calc()
}

func (d *day6) solvePart2() int64 {
	return d.calc2()
}

func (d *day6) calc() int64 {
	total := int64(0)

	for _, col := range d.columns {
		if col.operator == "+" {
			sum := int64(0)
			for _, v := range col.values {
				sum += v
			}

			total += sum
		} else if col.operator == "*" {
			prod := int64(1)
			for _, v := range col.values {
				prod *= v
			}

			total += prod
		}
	}

	return total
}

func (d *day6) calc2() int64 {
	total := int64(0)

	for _, col := range d.columns2 {
		if col.operator == "+" {
			sum := int64(0)
			for _, v := range col.values {
				sum += v
			}

			total += sum
		} else if col.operator == "*" {
			prod := int64(1)
			for _, v := range col.values {
				prod *= v
			}

			total += prod
		}
	}

	return total
}
