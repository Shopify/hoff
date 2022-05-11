package hoff

func Fill[T any](t T, num int) []T {
	out := make([]T, 0, num)
	for i := 0; i < num; i++ {
		out = append(out, t)
	}
	return out
}
