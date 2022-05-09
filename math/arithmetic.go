package math

import "github.com/akitasoftware/go-utils/constraints"

func Add[T constraints.Number](x, y T) T {
	return x + y
}

func Min[T constraints.Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Max[T constraints.Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}
