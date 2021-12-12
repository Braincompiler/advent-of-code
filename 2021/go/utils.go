package main

import (
	"fmt"
	"strconv"
)

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert %s to an integer number: %v", s, err))
	}

	return n
}

func max(a ...int) int {
	r := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > r {
			r = a[i]
		}
	}

	return r
}

func min(a ...int) int {
	r := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] < r {
			r = a[i]
		}
	}

	return r
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}
