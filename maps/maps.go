package maps

// ToValues returns an array of just the values of the input Map.
func ToValues[In any, Key comparable](arr map[Key]In) []In {
	i := 0
	out := make([]In, len(arr))
	for _, value := range arr {
		out[i] = value
		i++
	}
	return out
}

// ToSlice applies the function to each element of the input map, and returns an
// array of the results.
func ToSlice[M ~map[K]V, K comparable, V, R any](items M, f func(K, V) R) []R {
	i := 0
	out := make([]R, len(items))
	for k, v := range items {
		out[i] = f(k, v)
		i++
	}
	return out
}

func ToKeys[In any, Key comparable](arr map[Key]In) []Key {
	i := 0
	out := make([]Key, len(arr))
	for key := range arr {
		out[i] = key
		i++
	}
	return out
}
