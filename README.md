# BILLmanager 6 API for GoLang

## Install

> go get -u github.com/mayerdev/go-bm6-api

## Example

```go
package main

import "github.com/mayerdev/go-bm6-api"

func main() {
    api := bm6.New("https://bill.manager", "user", "password")

    api.Request(map[string]string{...})
}
```