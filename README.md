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
import "github.com/bettercode-oss/gin-middleware-etag/etag"
```
## Example
```go
package main

import (
"github.com/gin-gonic/gin"
"github.com/bettercode-oss/gin-middleware-etag/etag"
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