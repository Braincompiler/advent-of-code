package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type freshIdRange struct {
	start int64
	end   int64
}

type day5 struct {
	freshIds     []freshIdRange
	availableIds []int64
}

func Solve() {
	day := newDay5()
	day.parse()

	fmt.Println("===== Solution for Day 5 =====")
	fmt.Printf("    Part1: %d\n", day.solvePart1())
	fmt.Printf("    Part2: %d\n", day.solvePart2())
}

func newDay5() *day5 {
	return &day5{}
}

func (d *day5) parse() {
	file, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	isFreshRanges := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isFreshRanges = false
			continue
		}

		if isFreshRanges {
			var start, end int64
			fmt.Sscanf(line, "%d-%d", &start, &end)
			d.freshIds = append(d.freshIds, freshIdRange{start: start, end: end})
		} else {
			var id int64
			fmt.Sscanf(line, "%d", &id)
			d.availableIds = append(d.availableIds, id)
		}
	}
}

func (d *day5) solvePart1() int64 {
	numFresh := int64(0)
	for _, id := range d.availableIds {
		if isFresh(id, d.freshIds) {
			numFresh++
		}
	}

	return numFresh
}

func (d *day5) solvePart2() int64 {
	sort.Slice(d.freshIds, func(i, j int) bool {
		return d.freshIds[i].start < d.freshIds[j].start
	})

	merged := make([]freshIdRange, 1)
	merged[0] = d.freshIds[0]

	for i := 1; i < len(d.freshIds); i++ {
		var l = &merged[len(merged)-1]

		if d.freshIds[i].start <= l.end {
			l.end = max(l.end, d.freshIds[i].end)
		} else {
			merged = append(merged, d.freshIds[i])
		}
	}

	numTotalIds := int64(0)
	for _, r := range merged {
		numTotalIds += r.end - r.start + 1
	}

	return numTotalIds
}

func isFresh(id int64, freshRanges []freshIdRange) bool {
	for _, fr := range freshRanges {
		if id >= fr.start && id <= fr.end {
			return true
		}
	}

	return false
}
