package utils

// Chunk takes an input array of T elements, and "chunks" into groups of chunkSize elements.
// If the input array is empty, or the batchSize is 0, return an empty slice of slices.
// Adapted to generics from https://github.com/golang/go/wiki/SliceTricks#batching-with-minimal-allocation.
// Examples:
// Chunk([]int{1, 2, 3, 4, 5}, 2) = [[1, 2] [3, 4], [5]]
// Chunk([]int{}, 2) = []
// Chunk([]int{1, 2, 3}, 0) = [].
func Chunk[T any](actions []T, batchSize int) [][]T {
	// if the input is empty or batch size is 0, return an empty slice of slices
	if len(actions) == 0 || batchSize < 1 {
		return [][]T{}
	}
	// make out as a new slice of type T slices, up to the max number of chunks in the result
	batches := make([][]T, 0, (len(actions)+batchSize-1)/batchSize)

	for batchSize < len(actions) {
		actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
	}
	batches = append(batches, actions)
	return batches
}
