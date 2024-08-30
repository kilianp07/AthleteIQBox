package utils

import (
	"io"
	"log/syslog"
	"os"

	"github.com/rs/zerolog"
)

// Output represents a logging output with a writer and level.
type Output interface {
	GetWriter() (io.Writer, error)
	GetLevel() zerolog.Level
}

// ConsoleOutputConfig is the configuration for console output.
type ConsoleOutputConfig struct {
	Level string `json:"level"`
}

func (c ConsoleOutputConfig) ToOutput() Output {
	return ConsoleOutput{
		Level: ParseLogLevel(c.Level),
	}
}

// ConsoleOutput is a logging output for the console.
type ConsoleOutput struct {
	Level LogLevel
}

func (c ConsoleOutput) GetWriter() (io.Writer, error) {
	return zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}, nil
}

func (c ConsoleOutput) GetLevel() zerolog.Level {
	return c.Level.GetLevel()
}

// FileOutputConfig is the configuration for file output.
type FileOutputConfig struct {
	Level    string `json:"level"`
	Filename string `json:"filename"`
}

func (f FileOutputConfig) ToOutput() Output {
	return FileOutput{
		Level:    ParseLogLevel(f.Level),
		Filename: f.Filename,
	}
}

// FileOutput is a logging output to a file.
type FileOutput struct {
	Level    LogLevel
	Filename string
}

func (f FileOutput) GetWriter() (io.Writer, error) {
	return os.Create(f.Filename)
}

func (f FileOutput) GetLevel() zerolog.Level {
	return f.Level.GetLevel()
}

// SyslogOutputConfig is the configuration for syslog output.
type SyslogOutputConfig struct {
	Level string `json:"level"`
	Tag   string `json:"tag"`
}

func (s SyslogOutputConfig) ToOutput() Output {
	return SyslogOutput{
		Level: ParseLogLevel(s.Level),
		Tag:   s.Tag,
	}
}

// SyslogOutput is a logging output to the system's local syslog (/var/log/syslog).
type SyslogOutput struct {
	Level LogLevel
	Tag   string
}

func (s SyslogOutput) GetWriter() (io.Writer, error) {
	return syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, s.Tag)
}

func (s SyslogOutput) GetLevel() zerolog.Level {
	return s.Level.GetLevel()
}
