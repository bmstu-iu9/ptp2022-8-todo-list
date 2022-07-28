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
		logger.log.Print("[DEBUG] ")
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

// NewForTest creates a new logger using configuration for unit-tests.
func NewForTest() Logger {
	file, err := os.OpenFile("tests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return logger{
		log: log.New(file, "", log.Ldate|log.Ltime),
	}
}
