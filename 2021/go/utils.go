package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert %s to an integer number: %v", s, err))
	}

	return n
}

func parseUint64(s string) uint64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert %s to an integer number: %v", s, err))
	}

	return uint64(n)
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

func mapStringSlice(a []string, f func(string) string) []string {
	var result []string

	for _, s := range a {
		result = append(result, f(s))
	}

	return result
}

func contains(a []int, i int) bool {
	for _, n := range a {
		if n == i {
			return true
		}
	}

	return false
}

func sortString(s string) string {
	a := strings.Split(s, "")

	sort.Strings(a)

	return strings.Join(a, "")
}
