![Build Status](https://github.com/bettercode-oss/gin-middleware-etag/actions/workflows/build.yml/badge.svg)
[![codecov](https://codecov.io/gh/bettercode-oss/gin-middleware-etag/branch/main/graph/badge.svg?token=tNKcOjlxLo)](https://codecov.io/gh/bettercode-oss/gin-middleware-etag)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/bettercode-oss/gin-middleware-etag)

# HTTP ETag Cache Middleware
Gin middleware/handler to enable HTTP Etag support.

# Usage
## Start using it
Download and install it:
```shell
go get github.com/bettercode-oss/gin-middleware-etag
```
Import it in your code:
```go
import "github.com/bettercode-oss/gin-middleware-etag"
```
## Example
```go
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/bettercode-oss/gin-middleware-etag"
  "net/http"
)

func main() {
  r := gin.Default()
  r.GET("/products", etag.HttpEtagCache(120), getProducts)
  r.Run()
}

func getProducts(c *gin.Context) {
  products := []map[string]any{
    {
        "id":        1,
        "name":      "큰 잔",
        "listPrice": 1000,
    },
    {
        "id":        2,
        "name":      "작은 잔",
        "listPrice": 2000,
    },
  }

  c.JSON(http.StatusOK, products)
}
```
