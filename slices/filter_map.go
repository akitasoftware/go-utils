package slices

import "github.com/akitasoftware/go-utils/optionals"

// Applies f to each element of slice in order, removes any None results, and
// returns the rest.
func FilterMap[T1, T2 any](slice []T1, f func(T1) optionals.Optional[T2]) []T2 {
	result, _ := FilterMapIndexWithErr(slice, func(_ int, t T1) (optionals.Optional[T2], error) {
		return f(t), nil
	})
	return result
}

// Applies f to each element of slice in order, removes any None results, and
// returns the rest. If f returns a non-nil error on any element, iteration
// immediately stops, and the error is returned.
func FilterMapWithErr[T1, T2 any](slice []T1, f func(T1) (optionals.Optional[T2], error)) ([]T2, error) {
	return FilterMapIndexWithErr(slice, func(_ int, t T1) (optionals.Optional[T2], error) {
		return f(t)
	})
}

// Like FilterMap, but f also takes in the element's index.
func FilterMapIndex[T1, T2 any](slice []T1, f func(int, T1) optionals.Optional[T2]) []T2 {
	result, _ := FilterMapIndexWithErr(slice, func(idx int, t T1) (optionals.Optional[T2], error) {
		return f(idx, t), nil
	})
	return result
}

// Like FilterMapWithErr, but f also takes in the element's index.
func FilterMapIndexWithErr[T1, T2 any](slice []T1, f func(int, T1) (optionals.Optional[T2], error)) ([]T2, error) {
	if slice == nil {
		return nil, nil
	}

	result := make([]T2, 0, len(slice))
	for idx, t := range slice {
		u_opt, err := f(idx, t)
		if err != nil {
			return nil, err
		}
		if u, exists := u_opt.Get(); exists {
			result = append(result, u)
		}
	}

	return result, nil
}
