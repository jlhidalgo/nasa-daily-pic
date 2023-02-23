# Nasa's picture of the day

A simple project to get the NASA Astronomic Picture of the Day through its API. 

## Description

WIP

## Getting started

### Dependencies

This is required for building and compiling the application:
* Go version 1.18.x or higher

### Installing

1. Clone this repository.

2. Make the repository folder your working directory

WIP (further details may be added in upcoming weeks)

### Executing the program

1. In Linux and MacOS:

- You can use `make` to build the application:
```bash
    make build
```

- Then just run the app:
```bash
    ./bin/pic-server
```

2. In Windows:

- Build and run the application
```bash
    go build -o ./bin/ ./cmd/pic-server/pic-server.go
    ./bin/pic-server
```

### Testing

#### Unit tests

1. Linux and MacOS:
```bash
    make test
```

2. Windows:
```bash
    go test -v ./...
```

WIP (need to add directions on how to test the app from web browser, postman, etc.)

## License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details

## Acknowledgments
Inspiration, code snippets, etc.

* [NASA APIs](https://api.nasa.gov/)
* [Astronomy picture of the day](https://apod.nasa.gov/apod/astropix.html)
* [Testing with GoMock: A Tutorial](https://gist.github.com/thiagozs/4276432d12c2e5b152ea15b3f8b0012e)
* [gomock](https://github.com/golang/mock)
* [Testify](https://github.com/stretchr/testify)
* [Writing Web Applications](https://go.dev/doc/articles/wiki/)