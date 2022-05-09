# hoff: Higher Order Functions (and Friends)

Golang 1.18+ implementations of common methods/data structures using Go Generics

## Quick Start

[main.go](./main.go) contains a comprehensive example of how the `set` package works.

Compile and run the main.go file in the project root in one easy command:

```bash
go run ./main.go
```

## Running tests/benchmarks

Run the tests and benchmarks for the project using this command:

```bash
go test -v -bench=. ./...
```

## CI/CD and Github Actions

This project is configured to use GH Actions to automatically test/benchmark the project whenever pushes occur.
See the [.github/workflows](./.github/workflows) folder for all the details.

## License

hoff is released under the [MIT License](https://opensource.org/licenses/MIT).
