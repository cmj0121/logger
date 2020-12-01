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
	// [Logger] logger_test.go:9: show the CRIT log
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
	// [Logger] logger_test.go:21: show the CRIT log
	// [Logger] logger_test.go:22: show the WARN log
	// [Logger] logger_test.go:23: show the INFO log
	// [Logger] logger_test.go:24: show the DEBUG log
}

func ExampleCustomLog() {
	log := New("example")
	log.SetOutput(os.Stdout)
	log.SetLevel("info")
	log.Log(CRIT, "show the CRIT log")
	log.Log(WARN, "show the WARN log")
	log.Log(INFO, "show the INFO log")
	log.Log(DEBUG, "show the DEBUG log")
	log.Log(VERBOSE, "show the VERBOSE log")
	// Output:
	// [example] logger_test.go:37: show the CRIT log
	// [example] logger_test.go:38: show the WARN log
	// [example] logger_test.go:39: show the INFO log
}
