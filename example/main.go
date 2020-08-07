package main

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.New()

	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(true))

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example when panic happen.
	r.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
