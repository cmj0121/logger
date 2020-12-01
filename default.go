package logger

import (
	"os"
	"strings"
)

func init() {
	lv := strings.ToUpper(os.Getenv(ENV_LOG_LEVEL))
	default_log.SetLevel(lv)
	default_log.SetCallpath(3)
}

func Log(lv LogLevel, msg string, args ...interface{}) {
	default_log.Log(lv, msg, args...)
	return
}

func Crit(msg string, args ...interface{}) {
	default_log.Log(CRIT, msg, args...)
	return
}

func Warn(msg string, args ...interface{}) {
	default_log.Log(WARN, msg, args...)
	return
}

func Info(msg string, args ...interface{}) {
	default_log.Log(INFO, msg, args...)
	return
}

func Debug(msg string, args ...interface{}) {
	default_log.Log(DEBUG, msg, args...)
	return
}

func Verbose(msg string, args ...interface{}) {
	default_log.Log(VERBOSE, msg, args...)
	return
}
