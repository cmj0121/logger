package logger

import (
	"os"
)

func ExampleDefaultLog() {
	default_log.SetOutput(os.Stdout)
	Log(CRIT, "show the CRIT log")
	Log(WARN, "show the WARN log")
	Log(INFO, "show the INFO log")
	Log(DEBUG, "show the DEBUG log")
	Log(VERBOSE, "show the VERBOSE log")
	// Output:
	// [Logger] default_test.go:9: [CRIT] show the CRIT log
}

func ExampleDefaultLogLevel() {
	default_log.SetOutput(os.Stdout)
	default_log.SetLevel("debug")
	Log(CRIT, "show the CRIT log")
	Log(WARN, "show the WARN log")
	Log(INFO, "show the INFO log")
	Log(DEBUG, "show the DEBUG log")
	Log(VERBOSE, "show the VERBOSE log")
	// Output:
	// [Logger] default_test.go:21: [CRIT] show the CRIT log
	// [Logger] default_test.go:22: [WARN] show the WARN log
	// [Logger] default_test.go:23: [INFO] show the INFO log
	// [Logger] default_test.go:24: [DEBUG] show the DEBUG log
}
