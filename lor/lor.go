package lor

func Map[T any, R any](arr []T, f func(T) R) []R {
	result := make([]R, len(arr), cap(arr))

	for i, v := range arr {
		result[i] = f(v)
	}

	return result
}

func Filter[T any](arr []T, f func(T) bool) []T {
	result := make([]T, 0, len(arr))

	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce[T any, R any](arr []T, init R, f func(R, T) R) R {
	for _, v := range arr {
		init = f(init, v)
	}

	return init
}

func FlatMap[T any, R any](arr []T, f func(item T, index int) []R) []R {
	result := make([]R, 0, len(arr))

	for i, v := range arr {
		result = append(result, f(v, i)...)
	}

	return result
}
