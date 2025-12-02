package day1

import (
	"2024/utils"
	"sort"
	"strings"
)

type Day1 struct {
	leftNumbers  []int
	rightNumbers []int
}

func NewDay1() *Day1 {
	return &Day1{}
}

func (d *Day1) Parse(line string) {
	parts := strings.SplitN(line, " ", 2)

	d.leftNumbers = append(d.leftNumbers, utils.Atoi(strings.Trim(parts[0], " ")))
	d.rightNumbers = append(d.rightNumbers, utils.Atoi(strings.Trim(parts[1], " ")))
}

func (d *Day1) Solution1() int {
	sort.Ints(d.leftNumbers)
	sort.Ints(d.rightNumbers)

	sum := 0
	for i, leftNumber := range d.leftNumbers {
		rightNumber := d.rightNumbers[i]

		sum += utils.AbsInt(leftNumber - rightNumber)
	}

	return sum
}

func (d *Day1) Solution2() int {
	sum := 0
	for _, leftNumber := range d.leftNumbers {
		count := utils.Count(d.rightNumbers, func(i int) bool {
			return i == leftNumber
		})
		sum += leftNumber * count
	}

	return sum
}
