# hoff: Higher Order Functions (and Friends)

Golang 1.18+ implementations of common methods/data structures using Go Generics

## Requirements

- Go 1.18 or newer (must support Generics)

## In Development

Please note: this package is still under development and may change in the future. We will attempt to maintain as much
backwards compatibility as possible for future changes, but this is still a v0.x release and things might change.

Mash that Star button and OBLITERATE the Watch button to follow our changes.

## Running tests/benchmarks

Run the tests and benchmarks for the project using this command:

```bash
go test -v -bench=. -race ./...
```

## CI/CD and Github Actions

This project is configured to use GH Actions to automatically test/benchmark the project whenever pushes occur.
See the [.github/workflows](./.github/workflows) folder for all the details.

## Contributing

Contributors must sign the [Shopify CLA](https://cla.shopify.com/) before your PR can be accepted/merged.

## Authors

- [Chris Pappas](https://github.com/chrispappas)
- [Eduardo Cuducos](https://github.com/cuducos)
- [Jade Ornelas](https://github.com/yknx4)
- [Paco Castro](https://github.com/pacocastrotech)
- Please add a commit with however you want to be credited!!

## License

hoff is released under the [MIT License](https://opensource.org/licenses/MIT).
