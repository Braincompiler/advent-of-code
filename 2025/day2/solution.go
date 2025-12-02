package day2

import (
	"braincompiler/aoc2025/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	start int64
	end   int64
}

type day2 struct {
	idRanges []idRange
}

func Solve() {
	day := newDay2()
	day.parse()

	fmt.Println("===== Solution for Day 2 =====")
	fmt.Printf("    Part1: %d\n", day.solvePart1())
	fmt.Printf("    Part2: %d\n", day.solvePart2())
}

func newDay2() *day2 {
	return &day2{}
}

func (d *day2) parse() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitSeq(line, ",")

		for part := range parts {
			ids := strings.Split(part, "-")
			d.idRanges = append(d.idRanges, idRange{
				start: utils.Atoi64(ids[0]),
				end:   utils.Atoi64(ids[1]),
			})
		}
	}
}

func (d *day2) solvePart1() int64 {
	var invalidIds []int64

	for _, idRange := range d.idRanges {
		for n := range utils.RangeIncluded64(idRange.start, idRange.end) {
			id := strconv.FormatInt(n, 10)
			if len(id)%2 == 0 {
				mid := len(id) / 2
				left := id[:mid]
				right := id[mid:]

				if left == right {
					invalidIds = append(invalidIds, n)
				}
			}
		}
	}

	return utils.Sum64(invalidIds)
}

func (d *day2) solvePart2() int64 {
	var invalidIds []int64

	for _, idRange := range d.idRanges {
		for n := range utils.RangeIncluded64(idRange.start, idRange.end) {
			id := strconv.FormatInt(n, 10)
			idLen := int64(len(id))
			mid := idLen / 2
			for i := range utils.RangeIncluded64(1, int64(mid)) {
				if idLen%i != 0 {
					continue
				}

				pattern := id[:i]
				matches := true
				for j := i; j < idLen; j += i {
					if id[j:j+i] != pattern {
						matches = false
						break
					}
				}

				if matches {
					invalidIds = append(invalidIds, n)
					break
				}
			}
		}
	}

	return utils.Sum64(invalidIds)
}
