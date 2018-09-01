# go-exactonline #

[![GoDoc](https://godoc.org/github.com/mcnijman/go-exactonline?status.svg)](https://godoc.org/github.com/mcnijman/go-exactonline) [![Build Status](https://travis-ci.org/mcnijman/go-exactonline.svg?branch=master)](https://travis-ci.org/mcnijman/go-exactonline) [![Test Coverage](https://coveralls.io/repos/github/mcnijman/go-exactonline/badge.svg?branch=master)](https://coveralls.io/github/mcnijman/go-exactonline?branch=master)

go-exactonline is a Go client library for accessing the Exact Online API.

This is library is incomplete and under development.

## Usage ##

```go
import "github.com/mcnijman/go-exactonline"
```

Note that this library doens't directly handle authentication, see [Authentication](Authentication).

```go
client := exactonline.NewClient(nil)
```

## Authentication ##

This library doesn't directly handle authentication. You should provide a `http.Client` that handles the authentication for you.

## Divisions ##


## Issues ##

Issues and/or pull requests are welcome. Note that the services are generated using `gen-services.go`. If there are issues with the services, take a look at this file or the templates. The service files shouldn't be edites.

## Versioning ##

This library uses symantic versions using git tags. However since this library is still in development, API methods are subject to change.

## TODO ##

- Implement all CRUD operations
- Helpers for odata types creation
- Integration tests
- Error handling
- Web hooks support
- Documentation and examples

## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.