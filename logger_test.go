package logger

import (
	"os"
)

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
	// [example] logger_test.go:11: show the CRIT log
	// [example] logger_test.go:12: show the WARN log
	// [example] logger_test.go:13: show the INFO log
}

func Foo() {
	log := New("foo")
	log.SetOutput(os.Stdout)
	log.SetLevel("info")
	log.Log(CRIT, "show the CRIT log")
	log.Log(WARN, "show the WARN log")
	log.Log(INFO, "show the INFO log")
	log.Log(DEBUG, "show the DEBUG log")
	log.Log(VERBOSE, "show the VERBOSE log")
}

func ExampleFoo() {
	Foo()
	// Output:
	// [foo] logger_test.go:26: show the CRIT log
	// [foo] logger_test.go:27: show the WARN log
	// [foo] logger_test.go:28: show the INFO log
}
