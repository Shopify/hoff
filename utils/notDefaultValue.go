package utils

func NotDefaultValue[T comparable](t T) bool {
	var defaultValue T
	return t != defaultValue
}
