# RequestID

Request ID middleware for Gin Framework. Adds an indentifier to the response using the `X-Request-ID` header. Passes the `X-Request-ID` value back to the caller if it's sent in the request headers.

## Examples

### Using the middleware with the default config

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/printesoi/requestid"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	r.Use(requestid.New())

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```

### Custom generator function

```go
func main() {

	r := gin.New()

	r.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return "test"
		},
	}))

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```

### Getting the request identifier

```go
// Example / request.
r.GET("/", func(c *gin.Context) {
	c.String(http.StatusOK, "id:"+requestid.Get(c))
})
```

### Using a custom request ID header

```go
r.Use(requestid.New(requestid.Config{
	RequestIDHeader: "X-Custom-ID",
}))
```
