# hoff: Higher Order Functions (and Friends)

Golang 1.18+ implementations of common methods/data structures using Go Generics

## Requirements

- Go 1.18 or newer (must support Generics)

## Running tests/benchmarks

Run the tests and benchmarks for the project using this command:

```bash
go test -v -bench=. ./...
```

## CI/CD and Github Actions

This project is configured to use GH Actions to automatically test/benchmark the project whenever pushes occur.
See the [.github/workflows](./.github/workflows) folder for all the details.

## Contributing

Contributors must sign the [Shopify CLA](https://cla.shopify.com/) before your PR can be accepted/merged.

## License

hoff is released under the [MIT License](https://opensource.org/licenses/MIT).
