package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func WithTimeOut(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
