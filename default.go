package logger

import (
	"os"
	"strings"
)

func init() {
	lv := strings.ToUpper(os.Getenv(ENV_LOG_LEVEL))
	default_log.SetLevel(lv)
}

func Log(lv LogLevel, msg string, args ...interface{}) {
	default_log.calldepth = 3
	default_log.Log(lv, msg, args...)
	return
}
