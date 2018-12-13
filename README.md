go-jerakia
==========

A Go client library for [Jerakia](http://jerakia.io).

Quickstart
----------

```go
package main

import (
  "fmt"
  "net/http"

  "github.com/jerakia/go-jerakia"
)

func main() {
  config := jerakia.ClientConfig{
    URL:   "http://localhost:9992/v1",
    Token: "mytok:abcd",
  }

  client := jerakia.NewClient(http.DefaultClient, config)

  lookupOpts := jerakia.LookupOpts{
    Namespace: "test",
    Metadata: map[string]string{
      "hostname": "example",
    },
  }

  result, err := jerakia.Lookup(&client, "users", &lookupOpts)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%#v\n", result)
}
```

Development
-----------

### Installation

```shell
$ go get github.com/jerakia/go-jerakia
```

### Unit Tests

```shell
$ cd $GOPATH/github.com/jerakia/go-jerakia
$ make test
```

### Acceptance Tests

Make sure you have the following environment variables set:

* `JERAKIA_URL`
* `JERAKIA_TOKEN`

```shell
$ cd $GOPATH/github.com/jerakia/go-jerakia
$ make testacc
```

> You can use the supplied `acceptance/deploy.sh` script to install
> all requirements (including Jerakia and Go) on an Ubuntu 16.04 system.
>
> The script will create a `~/jrc` file with all required environment
> variables set.

### Vendor Dependencies

`go-jerakia` uses [Go modules](https://github.com/golang/go/wiki/Modules) for dependency/vendor management.
