package logger

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

type Config struct {
	Logger  *log.Logger
	Context log.Context
	Logging func(c *gin.Context) bool
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

		if newConfig.Logging != nil && !newConfig.Logging(c) {
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

		ctx := log.NewContext(nil).
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Str("user_agent", c.Request.UserAgent()).
			Value()

		switch {
		case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
			newConfig.Logger.Warn().Context(newConfig.Context).Context(ctx).Msg(msg)
		case c.Writer.Status() >= http.StatusInternalServerError:
			newConfig.Logger.Error().Context(newConfig.Context).Context(ctx).Msg(msg)
		default:
			newConfig.Logger.Info().Context(newConfig.Context).Context(ctx).Msg(msg)
		}
	}
}
