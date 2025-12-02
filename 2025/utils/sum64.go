package utils

func Sum64(ints []int64) int64 {
	sum := int64(0)
	for _, v := range ints {
		sum += v
	}

	return sum
}
