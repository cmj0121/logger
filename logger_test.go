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
	// [example] logger_test.go:11: [CRIT] show the CRIT log
	// [example] logger_test.go:12: [WARN] show the WARN log
	// [example] logger_test.go:13: [INFO] show the INFO log
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
	// [foo] logger_test.go:26: [CRIT] show the CRIT log
	// [foo] logger_test.go:27: [WARN] show the WARN log
	// [foo] logger_test.go:28: [INFO] show the INFO log
}
