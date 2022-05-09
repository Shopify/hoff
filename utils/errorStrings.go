package utils

// ErrorStrings is a convenience function to convert an array of errors to error strings.
func ErrorStrings(errs []error) []string {
	return Map(errs, func(err error) string {
		return err.Error()
	})
}
