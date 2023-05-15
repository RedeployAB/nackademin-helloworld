# nackademin-helloworld

> A simple hello world application for Nackademin's course DE22

* [Requirements](#requirements)
* [Test application](#test-application)
* [Build application](#build-application)
* [Build Dockerfile](#build-dockerfile)
* [Run application](#run-application)
    * [Run application with Go](#run-application-with-go)
    * [Run application from binary](#run-application-from-binary)
    * [Run application with Docker](#run-application-with-docker)

## Requirements

* [Go](https://go.dev/dl/) >= 1.20
* [Docker](https://docs.docker.com/get-docker/)


## Test application

```sh
go test ./...
```

## Build application

* `<os>` should contain the target OS, examples: `linux`, `darwin` and `windows`.
* `<arch>` should contain the target architecture, examples: `amd64` and `arm64`.


```
CGO_ENABLED=0 GOOS=<os> GOARCH=<arch> go build -o endpoint -o build/endpoint main.go
```

## Build Dockerfile

* `<version>` should contain a version number, example: `1.0.0`.

```sh
docker build -t endpoint:<version> .
```

## Run application

### Run application with Go

Run the application without compiling it with Go:
```sh
go run main.go
```

### Run application from binary

Run the application from the binary file (assuming it is built):
```sh
cd build
./endpoint

# Or
./build/endpoint
```

### Run application with Docker

* `<version>` should contain a version number, example: `1.0.0`. The same as used when building the Dockerfile.

```sh
docker run -d -p 8080:8080 --name endpoint endpoint:<version>
```
