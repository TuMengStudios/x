// Package ratelimit 备注, 代码来自于 github.com/go-kratos
package ratelimit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
)

// ErrLimitExceed is service unavailable due to rate limit exceeded.

// Option is ratelimit option.
type Option func(*options)

// WithLimiter set Limiter implementation,
// default is bbr limiter
func WithLimiter(limiter ratelimit.Limiter) Option {
	return func(o *options) {
		o.limiter = limiter
	}
}

type options struct {
	limiter ratelimit.Limiter
}

type Limit struct {
	limiter ratelimit.Limiter
}

func NewLimit(opts ...Option) *Limit {
	o := &options{
		limiter: bbr.NewLimiter(),
	}
	for _, opt := range opts {
		opt(o)
	}
	return &Limit{limiter: o.limiter}
}

func (c *Limit) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		done, err := c.limiter.Allow()
		if err != nil {
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		ctx.Next()
		// allowed
		done(ratelimit.DoneInfo{Err: err})
	}

}
