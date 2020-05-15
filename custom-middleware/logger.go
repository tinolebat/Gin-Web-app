package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s %s  %s - Latency: %s - Path: %s \n",
			param.ClientIP,
			param.Method,
			param.TimeStamp.Format(time.RFC822),
			param.Latency,
			param.Path,
		)
	})
}
