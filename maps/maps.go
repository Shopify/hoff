package maps

// ToValues returns an array of just the values of the input Map.
func ToValues[In any, Key comparable](arr map[Key]In) []In {
	out := make([]In, 0, len(arr))
	for _, elem := range arr {
		out = append(out, elem)
	}
	return out
}

// ToSlice applies the function to each element of the input map, and returns an
// array of the results.
func ToSlice[M ~map[K]V, K comparable, V, R any](items M, f func(K, V) R) []R {
	out := make([]R, 0, len(items))
	for k, v := range items {
		out = append(out, f(k, v))
	}
	return out
}
