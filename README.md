# reqgo

A simple wrapper for http requests

> Library is very much incomplete, its only purpose is for handling basic / normal http requests. Please use the default `net/http` package for more complex and control over the request

## Install

```sh
go get -u github.com/tbdsux/reqgo
```

## Usage

API is based from python's `requests` module and node's `fetch` api.

```go
package main

import (
	"fmt"
	"log"

	"github.com/tbdsux/reqgo"
)

func main() {
	r, err := reqgo.Get("https://httpbin.org/get", &reqgo.Options{
		Headers: reqgo.Headers{
			"Custom-Header": "hello world",
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	var data map[string]interface{}
	if err := r.JSON(&data); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data["headers"])
}

```

##

**&copy; 2023 | tbdsux**
