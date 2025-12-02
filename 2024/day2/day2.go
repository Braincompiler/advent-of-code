package day2

import (
	"2024/utils"
	"slices"
	"strings"
)

type LEVEL_DIRECTION int

const (
	NONE       LEVEL_DIRECTION = 0
	INCREASING LEVEL_DIRECTION = 1
	DECREASING LEVEL_DIRECTION = 2
	SAME       LEVEL_DIRECTION = 3
)

type Report struct {
	levels              []int
	directions          []LEVEL_DIRECTION
	directionsSet       *utils.Set[LEVEL_DIRECTION]
	differs             []int
	numDirectionChanges int
}

type Day2 struct {
	reports []*Report
}

func NewDay2() *Day2 {
	return &Day2{}
}

func NewReport(levels []int) *Report {
	return &Report{
		levels:        levels,
		directionsSet: utils.NewSet[LEVEL_DIRECTION](),
	}
}

func (r *Report) Calculate() {
	prevLevel := r.levels[0]
	prevDirection := NONE
	for levelIdx := 1; levelIdx < len(r.levels); levelIdx++ {
		level := r.levels[levelIdx]
		dir, differ := getLevelDirectionAndDiffer(level, prevLevel)

		if prevDirection != NONE && prevDirection != dir {
			r.numDirectionChanges += 1
		}

		r.directions = append(r.directions, dir)
		r.directionsSet.Add(dir)
		r.differs = append(r.differs, differ)

		prevLevel = level
		prevDirection = dir
	}

	slices.Sort(r.differs)
}

func (r *Report) NumDirections() int {
	return len(r.directions)
}

func (r *Report) NumUniqueDirections() int {
	return r.directionsSet.Size()
}

func (r *Report) NumDiffers() int {
	return len(r.differs)
}

func (r *Report) MaxDiffer() int {
	return slices.Max(r.differs)
}

func (r *Report) NumDirectionChanges() int {
	return r.numDirectionChanges
}

func (d *Day2) Parse(line string) {
	levelsAsStrings := strings.Split(line, " ")
	var levels []int
	for _, level := range levelsAsStrings {
		levels = append(levels, utils.Atoi(level))
	}

	report := NewReport(levels)
	report.Calculate()

	d.reports = append(d.reports, report)
}

func (d *Day2) Solution1() int {
	sum := 0

	for _, report := range d.reports {
		if report.NumUniqueDirections() == 1 && report.MaxDiffer() <= 3 {
			sum += 1
		} else {
			// fmt.Printf("Failed report: %v\n", report)
		}
	}

	return sum
}

func (d *Day2) Solution2() int {
	sum := 0

	for _, report := range d.reports {
		if report.NumDirectionChanges() <= 1 && report.MaxDiffer() <= 3 {
			sum += 1
		} else {
			// fmt.Printf("Failed report: %v\n", report)
		}
	}

	return sum
}

func getLevelDirectionAndDiffer(level, prevLevel int) (LEVEL_DIRECTION, int) {
	differ := utils.AbsInt(prevLevel - level)
	if level > prevLevel {
		return INCREASING, differ
	} else if level < prevLevel {
		return DECREASING, differ
	}

	return SAME, differ
}
