package utils

import (
	"fmt"
	"strings"
)

// KeyByField takes an array of structs and the field name,
// and returns a map keyed by the value of each item's fieldName.
// Note that a non-string-valued fieldName will cause errors.
func KeyByField[T any](items []T, fieldName string) (map[string]T, error) {
	keyed := make(map[string]T, len(items))
	for _, item := range items {
		key := GetStringField(item, fieldName)
		if strings.HasPrefix(key, "<") && strings.HasSuffix(key, "Value>") {
			// the key is not a string and thus will come back like <int value>
			return nil, fmt.Errorf("the field to key by is not a string")
		}
		keyed[key] = item
	}
	return keyed, nil
}
