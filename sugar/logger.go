package log

import (
	"github.com/phuslu/log"
)

// A Logger wraps the base Logger functionality in a slower, but less
// verbose, API. Any Logger can be converted to a Logger with its Sugar
// method.
//
// Unlike the Logger, the Logger doesn't insist on structured logging.
// For each log level, it exposes three methods: one for loosely-typed
// structured logging, one for println-style formatting, and one for
// printf-style formatting.
type Logger struct {
	Logger  log.Logger
	Context log.Context
}

// Level creates a child logger with the minimum accepted level set to level.
func (s *Logger) Level(level log.Level) *Logger {
	sl := *s
	if sl.Logger.Caller > 0 {
		sl.Logger.Caller += 1
	}
	sl.Logger.SetLevel(level)
	return &sl
}

// Print sends a log entry without extra field. Arguments are handled in the manner of fmt.Print.
func (s *Logger) Print(args ...interface{}) {
	s.Logger.Log().Context(s.Context).Msgs(args...)
}

// Println sends a log entry without extra field. Arguments are handled in the manner of fmt.Println.
func (s *Logger) Println(args ...interface{}) {
	s.Logger.Log().Context(s.Context).Msgs(args...)
}

// Printf sends a log entry without extra field. Arguments are handled in the manner of fmt.Printf.
func (s *Logger) Printf(format string, args ...interface{}) {
	s.Logger.Log().Context(s.Context).Msgf(format, args...)
}

// Log sends a log entry with extra fields.
func (s *Logger) Log(keysAndValues ...interface{}) error {
	s.Logger.Log().Context(s.Context).KeysAndValues(keysAndValues...).Msg("")
	return nil
}

// Debug uses fmt.Sprint to construct and log a message.
func (s *Logger) Debug(args ...interface{}) {
	s.Logger.Debug().Context(s.Context).Msgs(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (s *Logger) Debugf(template string, args ...interface{}) {
	s.Logger.Debug().Context(s.Context).Msgf(template, args...)
}

// Debugw logs a message with some additional context.
func (s *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	s.Logger.Debug().Context(s.Context).KeysAndValues(keysAndValues...).Msg(msg)
}

// Info uses fmt.Sprint to construct and log a message.
func (s *Logger) Info(args ...interface{}) {
	s.Logger.Info().Context(s.Context).Msgs(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (s *Logger) Infof(template string, args ...interface{}) {
	s.Logger.Debug().Context(s.Context).Msgf(template, args...)
}

// Infow logs a message with some additional context.
func (s *Logger) Infow(msg string, keysAndValues ...interface{}) {
	s.Logger.Info().Context(s.Context).KeysAndValues(keysAndValues...).Msg(msg)
}

// Warn uses fmt.Sprint to construct and log a message.
func (s *Logger) Warn(args ...interface{}) {
	s.Logger.Warn().Context(s.Context).Msgs(args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (s *Logger) Warnf(template string, args ...interface{}) {
	s.Logger.Warn().Context(s.Context).Msgf(template, args...)
}

// Warnw logs a message with some additional context.
func (s *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	s.Logger.Warn().Context(s.Context).KeysAndValues(keysAndValues...).Msg(msg)
}

// Error uses fmt.Sprint to construct and log a message.
func (s *Logger) Error(args ...interface{}) {
	s.Logger.Error().Context(s.Context).Msgs(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (s *Logger) Errorf(template string, args ...interface{}) {
	s.Logger.Error().Context(s.Context).Msgf(template, args...)
}

// Errorw logs a message with some additional context.
func (s *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	s.Logger.Error().Context(s.Context).KeysAndValues(keysAndValues...).Msg(msg)
}

// Fatal uses fmt.Sprint to construct and log a message.
func (s *Logger) Fatal(args ...interface{}) {
	s.Logger.Fatal().Context(s.Context).Msgs(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message.
func (s *Logger) Fatalf(template string, args ...interface{}) {
	s.Logger.Fatal().Context(s.Context).Msgf(template, args...)
}

// Fatalw logs a message with some additional context.
func (s *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	s.Logger.Fatal().Context(s.Context).KeysAndValues(keysAndValues...).Msg(msg)
}

// Panic uses fmt.Sprint to construct and log a message.
func (s *Logger) Panic(args ...interface{}) {
	s.Logger.Panic().Context(s.Context).Msgs(args...)
}

// Panicf uses fmt.Sprintf to log a templated message.
func (s *Logger) Panicf(template string, args ...interface{}) {
	s.Logger.Panic().Context(s.Context).Msgf(template, args...)
}

// Panicw logs a message with some additional context.
func (s *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	s.Logger.Panic().Context(s.Context).KeysAndValues(keysAndValues...).Msg(msg)
}
