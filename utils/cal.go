package utils

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}

type OrderedMax interface {
	int | float64
}

func Max[T OrderedMax](a, b T) T {
	if a > b {
		return a
	}
	return b
}
