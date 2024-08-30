package utils

import (
	"fmt"
	"os"

	"github.com/kilianp07/AthleteIQBox/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger manages logging to multiple outputs.
type Logger struct {
	conf    LoggerConfiguration
	loggers []zerolog.Logger
	prefix  string
}

var logger *Logger

func GetLogger(prefix string) *Logger {
	if logger == nil {
		return defaultLogger(prefix)
	}
	return logger.child(prefix)
}

func defaultLogger(prefix string) *Logger {
	defaultConfig := LoggerConfiguration{
		OutputsConfig: []map[string]interface{}{
			{
				"name":  "console",
				"level": "info",
			},
		},
		Timestamp: true,
		Caller:    true,
	}

	defaultLogger := &Logger{
		conf:    defaultConfig,
		loggers: []zerolog.Logger{zerolog.New(os.Stdout).With().Timestamp().Logger()},
		prefix:  prefix,
	}
	return defaultLogger
}

func (l *Logger) child(prefix string) *Logger {
	return &Logger{
		conf:    l.conf,
		loggers: l.loggers,
		prefix:  prefix,
	}
}

// NewLogger creates a new Logger instance.
func NewLogger(conf LoggerConfiguration) *Logger {
	logger = &Logger{
		conf:    conf,
		loggers: make([]zerolog.Logger, 0, len(conf.OutputsConfig)),
	}
	logger.configure()
	return logger
}

// configure sets up the loggers based on the configuration.
func (l *Logger) configure() {
	for _, outConfig := range l.conf.OutputsConfig {
		// Each output config is a map[string]interface{}, representing the raw JSON data for each output
		name, ok := outConfig["name"].(string)
		if !ok {
			log.Error().Msg("output 'name' is required and must be a string")
			continue
		}

		var output Output

		switch name {
		case "console":
			var consoleConfig ConsoleOutputConfig
			if err := decodeOutputConfig(outConfig, &consoleConfig); err != nil {
				log.Error().Err(err).Msg("failed to decode console output")
				continue
			}
			output = consoleConfig.ToOutput()

		case "file":
			var fileConfig FileOutputConfig
			if err := decodeOutputConfig(outConfig, &fileConfig); err != nil {
				log.Error().Err(err).Msg("failed to decode file output")
				continue
			}
			output = fileConfig.ToOutput()

		case "syslog":
			var syslogConfig SyslogOutputConfig
			if err := decodeOutputConfig(outConfig, &syslogConfig); err != nil {
				log.Error().Err(err).Msg("failed to decode syslog output")
				continue
			}
			output = syslogConfig.ToOutput()

		default:
			log.Error().Msgf("unknown output name: %s", name)
			continue
		}

		writer, err := output.GetWriter()
		if err != nil {
			log.Error().Err(err).Msgf("Failed to get writer for output: %s", name)
			continue
		}

		logger := zerolog.New(writer).Level(output.GetLevel())
		if l.conf.Timestamp {
			logger = logger.With().Timestamp().Logger()
		}
		if l.conf.Caller {
			logger = logger.With().Caller().Logger()
		}
		l.loggers = append(l.loggers, logger)
	}
}

// logMessage logs a message to all configured loggers.
func (l *Logger) logMessage(level zerolog.Level, msg string) {
	for _, logger := range l.loggers {
		if logger.GetLevel() <= level {
			logger.WithLevel(level).Msg(msg)
		}
	}
}

// Formatted logging methods for different levels.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logMessagef(zerolog.DebugLevel, format, args...)
}
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logMessagef(zerolog.InfoLevel, format, args...)
}
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logMessagef(zerolog.WarnLevel, format, args...)
}
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logMessagef(zerolog.ErrorLevel, format, args...)
}
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logMessagef(zerolog.FatalLevel, format, args...)
}

// logMessagef logs a formatted message to all configured loggers.
func (l *Logger) logMessagef(level zerolog.Level, format string, args ...interface{}) {
	if l.prefix != "" {
		format = l.prefix + ": " + format
	}
	msg := fmt.Sprintf(format, args...)
	l.logMessage(level, msg)
}

// decodeOutputConfig helps to decode the raw output configuration into the respective struct.
func decodeOutputConfig(input map[string]interface{}, output interface{}) error {
	decoder, err := utils.NewDecoder(output)
	if err != nil {
		return fmt.Errorf("failed to create decoder: %w", err)
	}
	if err := decoder.Decode(input); err != nil {
		return fmt.Errorf("failed to decode config: %w", err)
	}
	return nil
}
