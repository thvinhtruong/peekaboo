package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit(requestPerSec, maxRequestNum int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(requestPerSec), maxRequestNum)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
			return
		}

		c.Error(errors.New("number of requests exceeded"))
		c.AbortWithStatus(http.StatusTooManyRequests)
	}

}
