package utils

import "github.com/rs/zerolog"

// LogLevel represents a log level type with a GetLevel method.
type LogLevel interface {
	GetLevel() zerolog.Level
}

// LogLevel implementations for different log levels.
type (
	TraceLevel struct{}
	DebugLevel struct{}
	InfoLevel  struct{}
	WarnLevel  struct{}
	ErrorLevel struct{}
	FatalLevel struct{}
	PanicLevel struct{}
)

func (TraceLevel) GetLevel() zerolog.Level { return zerolog.TraceLevel }
func (DebugLevel) GetLevel() zerolog.Level { return zerolog.DebugLevel }
func (InfoLevel) GetLevel() zerolog.Level  { return zerolog.InfoLevel }
func (WarnLevel) GetLevel() zerolog.Level  { return zerolog.WarnLevel }
func (ErrorLevel) GetLevel() zerolog.Level { return zerolog.ErrorLevel }
func (FatalLevel) GetLevel() zerolog.Level { return zerolog.FatalLevel }
func (PanicLevel) GetLevel() zerolog.Level { return zerolog.PanicLevel }

// ParseLogLevel parses a string into a LogLevel implementation.
func ParseLogLevel(level string) LogLevel {
	switch level {
	case "trace":
		return TraceLevel{}
	case "debug":
		return DebugLevel{}
	case "info":
		return InfoLevel{}
	case "warn":
		return WarnLevel{}
	case "error":
		return ErrorLevel{}
	case "fatal":
		return FatalLevel{}
	case "panic":
		return PanicLevel{}
	default:
		return InfoLevel{}
	}
}
