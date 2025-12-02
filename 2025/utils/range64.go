package utils

func Range64(start, end int64) chan int64 {
	ch := make(chan int64)

	go func() {
		for i := start; i < end; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func RangeIncluded64(start, end int64) chan int64 {
	ch := make(chan int64)

	go func() {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}
