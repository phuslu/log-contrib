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
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.Logger == nil {
		cfg.Logger = &log.DefaultLogger
	}

	return func(c *fiber.Ctx) error {
		start := time.Now()
		next := c.Next()

		if cfg.Skip != nil && cfg.Skip(c) {
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
			e = cfg.Logger.Warn()
		case status >= 500:
			e = cfg.Logger.Error()
		default:
			e = cfg.Logger.Info()
		}
		e.Context(cfg.Context).
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
