package log

import "github.com/rs/zerolog"

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelPanic
)

type Level zerolog.Level

func (lvl *Level) Zerolog() zerolog.Level {
	return zerolog.Level(*lvl)
}
