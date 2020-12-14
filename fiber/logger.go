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

		ctx := log.NewContext(nil).
			Int("status", status).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("ip", c.IP()).
			Dur("latency", latency).
			Str("user_agent", c.Get(fiber.HeaderUserAgent)).
			Value()

		switch {
		case status >= 400 && status < 500:
			newConfig.Logger.Warn().Context(newConfig.Context).Context(ctx).Msg(msg)
		case status >= 500:
			newConfig.Logger.Error().Context(newConfig.Context).Context(ctx).Msg(msg)
		default:
			newConfig.Logger.Info().Context(newConfig.Context).Context(ctx).Msg(msg)
		}

		return nil
	}
}
