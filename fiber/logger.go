package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
)

type Config struct {
	Logger  *log.Logger
	Context log.Context
	Skip    func(c *fiber.Ctx) bool
}

func SetLogger(config ...Config) fiber.Handler {
	var newConfig Config
	if len(config) > 0 {
		newConfig = config[0]
	}

	if newConfig.Logger == nil {
		newConfig.Logger = &log.DefaultLogger
	}

	return func(c *fiber.Ctx) error {
		start := time.Now()
		next := c.Next()

		if newConfig.Skip != nil && !newConfig.Skip(c) {
			return nil
		}

		end := time.Now()
		latency := end.Sub(start)

		status := c.Response().StatusCode()
		msg := "Request"
		if next != nil {
			msg = next.Error()
		}

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
			Int("status", status).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("ip", c.IP()).
			Dur("latency", latency).
			Str("user_agent", c.Get(fiber.HeaderUserAgent)).
			Msg(msg)

		return nil
	}
}
