package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

const skipFrameCount = 3

type Logger interface {
	Debug(msg any, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg any, args ...any)
	Fatal(msg any, args ...any)
}

type Log struct {
	logger *zerolog.Logger
}

var _ Logger = (*Log)(nil)

func New(level string) *Log {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(l)

	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Log{
		logger: &logger,
	}
}

func (l *Log) log(msg string, args ...any) {
	if len(args) == 0 {
		l.logger.Info().Msg(msg)
	} else {
		l.logger.Info().Msgf(msg, args...)
	}
}

func (l *Log) msg(level string, message any, args ...any) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}

func (l *Log) Debug(msg any, args ...any) {
	l.msg("debug", msg, args...)
}

func (l *Log) Info(msg string, args ...any) {
	l.msg("info", msg, args...)
}

func (l *Log) Warn(msg string, args ...any) {
	l.msg("warn", msg, args...)
}

func (l *Log) Error(msg any, args ...any) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(msg, args...)
	}

	l.msg("error", msg, args...)
}

func (l *Log) Fatal(msg any, args ...any) {
	l.msg("fatal", msg, args...)

	os.Exit(1)
}
