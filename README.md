[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/sfpyhub/go-sfpy)

## Go SFPY

The official [SFPY](https://sfpy.co) Go client library.

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its root if it already is):

``` sh
go mod init
```

Then, reference go-sfpy in a Go program with `import`:

``` go
import (
    "github.com/sfpyhub/go-sfpy/sfpy"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the go-sfpy module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```
go get -u "github.com/sfpyhub/go-sfpy/sfpy"
```

## Initialize client

To use the SFPY client to make API requests, initialize it like so from within your service

```go
client := sfpy.NewClient(apikey, secretkey)
```

`NewClient` accepts two arguments:
1. `apikey`: This is your SFPY Api key that is needed to authenticate all requests to the API. You can find this key on your dashboard.
2. `secretkey`: This is your webhook shared secret key used to validate whether incoming signatures from webhooks are valid.