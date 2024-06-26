package log

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func Init(lvl Level) {
	zerolog.CallerSkipFrameCount = 3
	zerolog.LevelFieldName = "lvl"
	zerolog.ErrorFieldName = "err"
	zerolog.TimestampFieldName = "dt"
	zerolog.CallerFieldName = "call"
	zerolog.MessageFieldName = "msg"

	logger = zerolog.New(os.Stdout).
		Level(lvl.Zerolog()).
		With().Timestamp().Caller().
		Logger()
}

func AddHook(hook zerolog.Hook) {
	logger = logger.Hook(hook)
}

func Trace(msg string) {
	defer recovery()
	logger.Trace().Msg(msg)
}

func Tracef(format string, args ...interface{}) {
	defer recovery()
	logger.Trace().Msgf(format, args...)
}

func Debug(msg string) {
	defer recovery()
	logger.Debug().Msg(msg)
}

func Debugf(format string, args ...interface{}) {
	defer recovery()
	logger.Debug().Msgf(format, args...)
}

func Info(msg string) {
	defer recovery()
	logger.Info().Msg(msg)
}

func Infof(format string, args ...interface{}) {
	defer recovery()
	logger.Info().Msgf(format, args...)
}

func Warning(msg string) {
	defer recovery()
	logger.Warn().Msg(msg)
}

func Warningf(format string, args ...interface{}) {
	defer recovery()
	logger.Warn().Msgf(format, args...)
}

func Error(msg string, err error) {
	defer recovery()
	logger.Error().Err(err).Msg(msg)
}

func Errorf(format string, args ...interface{}) {
	defer recovery()
	logger.Error().Msgf(format, args...)
}

func Fatal(msg string, err error) {
	logger.Fatal().Err(err).Msg(msg)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatal().Msgf(format, args...)
}

func Panic(msg string, err error) {
	logger.Panic().Err(err).Msg(msg)
}

func Panicf(format string, args ...interface{}) {
	logger.Panic().Msgf(format, args...)
}

func recovery() {
	if data := recover(); data != nil {
		Errorf("RECOVERED: %v", data)
	}
}
