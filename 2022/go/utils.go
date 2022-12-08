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

func max[T constraints.Integer | constraints.Float](a ...T) T {
	r := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > r {
			r = a[i]
		}
	}

	return r
}

func min[T constraints.Integer | constraints.Float](a ...T) T {
	r := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] < r {
			r = a[i]
		}
	}

	return r
}

func abs[T constraints.Integer | constraints.Float](a T, multiplyBy T) T {
	if a < 0 {
		return a * multiplyBy
	}

	return a
}

func sum[T constraints.Integer | constraints.Float](a ...T) T {
	var sum T

	for _, n := range a {
		sum += n
	}

	return sum
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

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

type NumberType interface {
	constraints.Integer | constraints.Float
}

func IsBetween[T NumberType](v, lo, hi T) bool {
	return v >= lo && v <= hi
}

type (
	Stack[T any] struct {
		top    *node[T]
		length int
	}

	node[T any] struct {
		value *T
		prev  *node[T]
	}
)

// NewStack Create a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

// Len the number of items in the stack
func (s *Stack[T]) Len() int {
	return s.length
}

// Peek the top item on the stack
func (s *Stack[T]) Peek() *T {
	if s.length == 0 {
		return nil
	}

	return s.top.value
}

// Pop the item of the stack and return it
func (s *Stack[T]) Pop() *T {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--

	return n.value
}

// PopMany items of the stack and return it
func (s *Stack[T]) PopMany(num int) []*T {
	if s.length == 0 {
		return nil
	}

	var values []*T
	for i := 0; i < num; i++ {
		n := s.top
		s.top = n.prev
		s.length--

		values = append(values, n.value)
	}

	return values
}

// Push a value onto the top of the stack
func (s *Stack[T]) Push(value T) *Stack[T] {
	n := &node[T]{&value, s.top}
	s.top = n
	s.length++

	return s
}

// PushMany values onto the top of the stack
func (s *Stack[T]) PushMany(values []*T) *Stack[T] {
	for _, value := range values {
		s.Push(*value)
	}

	return s
}

func sortMapByValues[TKey comparable, TValue constraints.Ordered](m map[TKey]TValue) {
	keys := make([]TKey, 0, len(m))
	newMap := make(map[TKey]TValue, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	for _, k := range keys {
		//fmt.Println(k, m[k])
		newMap[k] = m[k]
	}
}
