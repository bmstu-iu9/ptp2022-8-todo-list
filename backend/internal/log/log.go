package log

import (
	"log"
	"os"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/config"
)

var isDebug bool

func init() {
	if config.Get("RUNTIME_MODE") == "debug" {
		isDebug = true
	}
}

// Logger is a logger with simple interface. It supports two log levels:
// Debug and Release. On Release mode messages will be printed only with
// Info level.
type Logger interface {
	// Debug logs a message with Debug level.
	Debug(v ...interface{})
	// Info logs a message with Info level.
	Info(v ...interface{})
}

type logger struct {
	log *log.Logger
}

func (logger logger) Debug(v ...interface{}) {
	if isDebug {
		logger.log.Println(v...)
	}
}

func (logger logger) Info(v ...interface{}) {
	logger.log.Println(v...)
}

// New creates a new logger using the default configuration.
func New() Logger {
	return logger{
		log: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}
