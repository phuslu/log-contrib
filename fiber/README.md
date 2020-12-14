## fiber

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	fiberlogger "github.com/phuslu/log-contrib/fiber"
)

func main() {
	if log.IsTerminal(os.Stderr.Fd()) {
		log.DefaultLogger = log.Logger{
			TimeFormat: "15:04:05",
			Caller:     1,
			Writer: &log.ConsoleWriter{
				ColorOutput:    true,
				QuoteString:    true,
				EndWithMessage: true,
			},
		}
	}

	app := fiber.New()

	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	app.Use(fiberlogger.SetLogger())

	// Custom logger
	app.Use(fiberlogger.SetLogger(fiberlogger.Config{
		Logger: &log.Logger{
			Writer: &log.FileWriter{
				Filename: "access.log",
				MaxSize:  1024 * 1024 * 1024,
			},
		},
		Context: log.NewContext(nil).Str("foo", "bar").Value(),
		Skip: func(c *fiber.Ctx) bool {
			if string(c.Path()) == "/backdoor" {
				return true
			}
			return false
		},
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong " + fmt.Sprint(time.Now().Unix()))
	})

	app.Get("/backdoor", func(c *fiber.Ctx) error {
		return c.SendString("a backdoor, go away")
	})

	log.Fatal().Err(app.Listen(":3000")).Msg("")
}
```
