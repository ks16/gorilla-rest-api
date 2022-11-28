package log

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Logger struct {
	*zerolog.Logger
}

func NewLogger(name string, debug bool) *Logger {

	// log time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// logger creation
	logger := log.Output(os.Stdout).With().Str("service", name).Logger()

	// global logger setup
	log.Logger = logger

	return &Logger{&logger}
}

func Ctx(ctx context.Context) *Logger {
	return &Logger{zerolog.Ctx(ctx)}
}
