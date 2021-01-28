package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	headerXRequestID   = "X-Request-ID"
	requestIDHeaderKey = "requestIDHeaderKey"
)

// Config defines the config for RequestID middleware
type Config struct {
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	Generator func() string

	// RequestIdHeader defines a custom request ID header name instead of the
	// default X-Request-ID. Optional.
	RequestIdHeader string
}

// New initializes the RequestID middleware.
func New(config ...Config) gin.HandlerFunc {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}

	// Set config default values
	if cfg.Generator == nil {
		cfg.Generator = func() string {
			return uuid.New().String()
		}
	}
	if cfg.RequestIdHeader == "" {
		cfg.RequestIdHeader = headerXRequestID
	}

	return func(c *gin.Context) {
		c.Set(requestIDHeaderKey, cfg.RequestIdHeader)

		// Get id from request
		rid := c.GetHeader(cfg.RequestIdHeader)
		if rid == "" {
			rid = cfg.Generator()
		}

		// Set the id to ensure that the requestid is in the response
		c.Header(cfg.RequestIdHeader, rid)
		c.Next()
	}
}

// Get returns the request identifier
func Get(c *gin.Context) string {
	requestIDHeader := c.GetString(requestIDHeaderKey)
	return c.Writer.Header().Get(requestIDHeader)
}
