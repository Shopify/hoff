package utils

func Filter[T any](arr []T, fn func(T) bool) (out []T) {
	for _, elem := range arr {
		if fn(elem) {
			out = append(out, elem)
		}
	}
	return out
}
