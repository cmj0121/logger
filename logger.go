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

func (lv LogLevel) String() (str string) {
	levels := []string{
		"CRIT",
		"WARN",
		"INFO",
		"DEBUG",
		"VERBOSE",
	}

	if str = levels[VERBOSE]; lv >= CRIT && lv <= VERBOSE {
		str = levels[lv]
	}
	return
}

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

	// change the log level from ENV
	env := strings.ToUpper(fmt.Sprintf("%v_%v", name, ENV_LOG_LEVEL))
	lv := strings.ToUpper(os.Getenv(env))
	logger.SetLevel(lv)
	return
}

func GetLogger(name string) (logger *Logger) {
	logger, _ = logger_pool[name]
	return
}

func (logger *Logger) SetCallpath(calldepth int) {
	logger.Lock()
	defer logger.Unlock()

	logger.calldepth = calldepth
	return
}

func (logger *Logger) SetLevel(lv string) {
	if level, ok := log_level[strings.ToUpper(lv)]; ok {
		// override the log level by ENV
		logger.LogLevel = level
	}
}

func (logger *Logger) log(calldepth int, lv LogLevel, msg string, args ...interface{}) {
	logger.Lock()
	defer logger.Unlock()

	if lv <= logger.LogLevel {
		msg := fmt.Sprintf(msg, args...)
		msg = fmt.Sprintf("[%v] %v", lv, msg)
		logger.Logger.Output(calldepth, msg)
	}
}

func (logger *Logger) Log(lv LogLevel, msg string, args ...interface{}) {
	logger.log(logger.calldepth+1, lv, msg, args...)
	return
}

func (logger *Logger) Crit(msg string, args ...interface{}) {
	logger.log(logger.calldepth+1, CRIT, msg, args...)
	return
}

func (logger *Logger) Warn(msg string, args ...interface{}) {
	logger.log(logger.calldepth+1, WARN, msg, args...)
	return
}

func (logger *Logger) Info(msg string, args ...interface{}) {
	logger.log(logger.calldepth+1, INFO, msg, args...)
	return
}

func (logger *Logger) Debug(msg string, args ...interface{}) {
	logger.log(logger.calldepth+1, DEBUG, msg, args...)
	return
}

func (logger *Logger) Verbose(msg string, args ...interface{}) {
	logger.log(logger.calldepth+1, VERBOSE, msg, args...)
	return
}
