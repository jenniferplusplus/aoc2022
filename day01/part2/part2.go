package part2

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"sort"
)

type ClosedOrderedQueue[T constraints.Ordered] struct {
	elements []T
}

func (q *ClosedOrderedQueue[T]) Init(size int64) {
	q.elements = make([]T, 0, size)
}

func (q *ClosedOrderedQueue[T]) Push(element T) int {
	//q.sort()
	i, _ := slices.BinarySearch(q.elements, element)
	if i == 0 && len(q.elements) > 0 {
		return len(q.elements)
	}

	if len(q.elements) == cap(q.elements) {
		q.Pop()
		i--
	}
	q.elements = slices.Insert(q.elements, i, element)

	return len(q.elements)
}

func (q *ClosedOrderedQueue[T]) Pop() (value T, found bool) {
	l := len(q.elements)
	if l == 0 {
		var result T
		return result, false
	}

	last := q.elements[0]
	q.elements = q.elements[1:]

	return last, true
}

func (q *ClosedOrderedQueue[T]) sort() {
	if !sort.SliceIsSorted(q.elements, func(i, j int) bool { return less(i, j, q.elements) }) {
		sort.Slice(q.elements, func(i, j int) bool { return less(i, j, q.elements) })
	}
}

func less[T constraints.Ordered](i, j int, list []T) bool {
	return list[i] < list[j]
}

func insertionPoint[T constraints.Ordered](list []T, low int, high int, value T) (index int) {
	var mid int = low + (high-low)/2
	switch {
	case mid >= len(list):
		return mid
	case low > high:
		return low
	case list[mid] == value:
		return insertionPoint(list, low, mid-1, value)
	case list[mid] > value:
		return insertionPoint(list, low, mid-1, value)
	default:
		return insertionPoint(list, mid+1, high, value)
	}
}
