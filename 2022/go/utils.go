package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
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

func containsString(a []string, i string) bool {
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

//func permutations(a []int) [][]int {
//	var result [][]int
//
//	perm(a, func(ints []int) {
//		result = append(result, ints)
//	}, 0)
//
//	return result
//}
//
//// Permute the values at index i to len(a)-1.
//func perm(a []int, f func([]int), i int) {
//	if i > len(a) {
//		f(a)
//		return
//	}
//
//	perm(a, f, i+1)
//
//	for j := i + 1; j < len(a); j++ {
//		a[i], a[j] = a[j], a[i]
//		perm(a, f, i+1)
//		a[i], a[j] = a[j], a[i]
//	}
//}

func permutations(xs []int) (permuts [][]int) {
	var rc func([]int, int)

	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}

	rc(xs, 0)

	return permuts
}

type ZipPair struct {
	A rune
	B int
}

func zip(s string, a []int) []ZipPair {
	var zipPairs []ZipPair

	for i := 0; i < len(s); i++ {
		zipPairs = append(zipPairs, ZipPair{
			A: rune(s[i]),
			B: a[i],
		})
	}

	return zipPairs
}

type NumberType interface {
	constraints.Integer | constraints.Float
}

func IsBetween[T NumberType](v, lo, hi T) bool {
	return v >= lo && v <= hi
}
