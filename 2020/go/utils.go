package main

import (
	"fmt"
	"strconv"
)

func parseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert %s to an integer number", s))
	}

	return num
}
