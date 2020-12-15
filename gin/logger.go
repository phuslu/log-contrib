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
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.Logger == nil {
		cfg.Logger = &log.DefaultLogger
	}

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		if cfg.Skip != nil && cfg.Skip(c) {
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
			e = cfg.Logger.Warn()
		case status >= 500:
			e = cfg.Logger.Error()
		default:
			e = cfg.Logger.Info()
		}
		e.Context(cfg.Context).
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Str("user_agent", c.Request.UserAgent()).
			Msg(msg)
	}
}
