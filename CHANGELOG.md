# Changelog

## [0.2.0] - 2022-05-17

[Refactor Result Values, Errors and add Error method](https://github.com/Shopify/hoff/pull/27)

- Update Result.Values to only return values for non-errored Results
- Update Result.Errors to only return errors for errored Results
- Add Result.Error which combines, comma-separated, all the error messages into a single error

## [0.1.0] - 2022-05-11

### Initial Release

See https://pkg.go.dev/github.com/Shopify/hoff

This library contains Generics implementations of the following types of functions to work on simple arrays/slices:

- chunk
- fill
- filter
- flatmap
- foreach
- map
- pluck
- reduce

Many of these methods come with variants capable of passing a context through, returning errors in addition to their
result, and both context and errors.

As well as functions to work on maps:

- ToValues
- ToSlice
