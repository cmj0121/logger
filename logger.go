package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// the log level
type LogLevel int

const (
	CRIT LogLevel = iota
	WARN
	INFO
	DEBUG
	VERBOSE
)

const (
	ENV_LOG_LEVEL = "LOG_LEVEL"
	PROJ_NAME     = "Logger"

	MAJOR = 1
	MINOR = 0
	MACRO = 0
)

var (
	logger_lock = sync.Mutex{}
	logger_pool = map[string]*Logger{}
	default_log = New(PROJ_NAME)
)

type Logger struct {
	// control the log level
	LogLevel

	// the lower-end logger system
	*log.Logger `-`

	// lock
	sync.Mutex `-`

	// the call path
	calldepth int `-`
}

func New(name string) (logger *Logger) {
	logger_lock.Lock()
	defer logger_lock.Unlock()

	logger = &Logger{
		LogLevel:  CRIT,
		Logger:    log.New(os.Stderr, fmt.Sprintf("[%v] ", name), log.Lshortfile),
		calldepth: 2,
	}

	logger_pool[name] = logger
	return
}

func (logger *Logger) Log(lv LogLevel, msg string, args ...interface{}) {
	logger.Lock()
	defer logger.Unlock()

	if lv <= logger.LogLevel {
		logger.Logger.Output(logger.calldepth, fmt.Sprintf(msg, args...))
	}
}

func (logger *Logger) SetLevel(lv string) {
	logs := map[string]LogLevel{
		"CRIT":    CRIT,
		"WARN":    WARN,
		"INFO":    INFO,
		"DEBUG":   DEBUG,
		"VERBOSE": VERBOSE,
	}

	if level, ok := logs[strings.ToUpper(lv)]; ok {
		// override the log level by ENV
		logger.LogLevel = level
	}
}

func init() {
	lv := strings.ToUpper(os.Getenv(ENV_LOG_LEVEL))
	default_log.SetLevel(lv)
}

func Log(lv LogLevel, msg string, args ...interface{}) {
	default_log.calldepth = 3
	default_log.Log(lv, msg, args...)
	return
}