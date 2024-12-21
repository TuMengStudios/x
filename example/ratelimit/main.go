package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/aegis/ratelimit/bbr"

	"github.com/TuMengStudios/x/ratelimit"
)

func main() {
	limit := ratelimit.NewLimit(ratelimit.WithLimiter(bbr.NewLimiter()))
	var engine = gin.New()
	engine.Use(limit.Handler())
}
