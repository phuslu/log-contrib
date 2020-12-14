## gin

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	ginlogger "github.com/phuslu/log-contrib/gin"
	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
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

	if gin.IsDebugging() {
		log.DefaultLogger.SetLevel(log.DebugLevel)
	}

	r := gin.New()

	// Add a gin logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	r.Use(ginlogger.SetLogger())

	// Custom gin logger
	r.Use(ginlogger.SetLogger(ginlogger.Config{
		Logger: &log.Logger{
			Writer: &log.FileWriter{
				Filename: "access.log",
				MaxSize:  1024 * 1024 * 1024,
			},
		},
		Context: log.NewContext(nil).Str("foo", "bar").Value(),
		Logging: func(c *gin.Context) bool {
			if c.Request.URL.Path == "/backdoor" {
				return true
			}
			return false
		},
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	r.GET("/backdoor", func(c *gin.Context) {
		c.String(http.StatusOK, "a backdoor, go away")
	})

	r.Run(":8080")
}

```
