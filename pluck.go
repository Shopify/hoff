package main

// Pluck takes an input array of maps with K keys and V values and "Plucks" the selected keys into an array of arrays of V values.
// If the input array is empty, or the keys is 0, return an empty slice of slices.
func Pluck[M ~map[K]V, K comparable, V any](maps []M, keys ...K) [][]V {
	if len(keys) == 0 || len(maps) == 0 {
		return [][]V{}
	}
	result := make([][]V, 0, len(maps))

	for _, m := range maps {
		content := make([]V, 0, len(keys))
		for _, k := range keys {
			content = append(content, m[k])
		}
		result = append(result, content)
	}

	return result
}
