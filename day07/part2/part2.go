package part2

import (
	"errors"
	"sort"
)

func First(enumerable []int64, fn func(int64) bool) (int64, error) {
	for _, t := range enumerable {
		if fn(t) {
			return t, nil
		}

	}
	return 0, errors.New("not found")
}

func Sort(enumerable chan int64) []int64 {
	out := make([]int64, 0, 2)
	for each := range enumerable {
		out = append(out, each)
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i] <= out[j]
	})

	return out
}

func Filter(values chan int64, fn func(int64) bool) chan int64 {
	out := make(chan int64)
	go func() {
		defer close(out)
		for v := range values {
			if fn(v) {
				out <- v
			}
		}
	}()
	return out
}
