package order

import "sort"

type Ordered interface {
	~int | ~float64 | ~string
}

type OrderedSlice[T Ordered] []T

func (s OrderedSlice[T]) Len() int {
	return len(s)
}

func (s OrderedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s OrderedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SortType[T any] struct {
	slice   []T
	compare func(T, T) bool
}

func (s SortType[T]) Len() int {
	return len(s.slice)
}

func (s SortType[T]) Less(i, j int) bool {
	return s.compare(s.slice[i], s.slice[j])
}

func (s SortType[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func PerformSort[T any](slice []T, compare func(T, T) bool) {
	sort.Sort(SortType[T]{slice, compare})
}

func GenricMap[T1, T2 any](input []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func GenericFilter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, val := range input {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}
