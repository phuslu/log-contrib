package gin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

type Config struct {
	Logger  *log.Logger
	Context log.Context
	Skip    func(c *gin.Context) bool
}

func SetLogger(config ...Config) gin.HandlerFunc {
	var newConfig Config
	if len(config) > 0 {
		newConfig = config[0]
	}

	if newConfig.Logger == nil {
		newConfig.Logger = &log.DefaultLogger
	}

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		if newConfig.Skip != nil && !newConfig.Skip(c) {
			return
		}

		end := time.Now()
		latency := end.Sub(start)

		path := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			path = path + "?" + c.Request.URL.RawQuery
		}
		msg := "Request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}
		status := c.Writer.Status()

		var e *log.Entry
		switch {
		case status >= 400 && status < 500:
			e = newConfig.Logger.Warn()
		case status >= 500:
			e = newConfig.Logger.Error()
		default:
			e = newConfig.Logger.Info()
		}
		e.Context(newConfig.Context).
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Str("user_agent", c.Request.UserAgent()).
			Msg(msg)
	}
}
