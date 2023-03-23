package api

import (
	"github.com/gin-gonic/gin"
)

// This is just a placeholder for any middleware you might need.
func SampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add your middleware logic here.

		c.Next()
	}
}
