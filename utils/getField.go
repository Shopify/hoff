package utils

import "reflect"

func GetStringField[T any](t T, prop string) string {
	value := reflect.Indirect(reflect.ValueOf(t))
	if value.Kind() != reflect.Struct {
		return ""
	}
	return value.FieldByName(prop).String()
}
