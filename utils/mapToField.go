package utils

func MapToStringField[T any](tt []T, prop string) []string {
	out := make([]string, 0, len(tt))
	for _, elem := range tt {
		out = append(out, GetStringField(elem, prop))
	}
	return out
}
