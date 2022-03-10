# couch - [![Go Reference](https://pkg.go.dev/badge/github.com/anoriqq/couch.svg)](https://pkg.go.dev/github.com/anoriqq/couch)

Simple full text search engine.

## Usage

```shell
go get github.com/anoriqq/couch
```

example project is [here](https://github.com/anoriqq/couch-example).

## Development

```shell
# install requirements
go install github.com/tenntenn/testtime@latest
go mod tidy

# run all test
go test -race -overlay=$(testtime) -shuffle on -count 1 ./...
```

