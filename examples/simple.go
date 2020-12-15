package main

import (
	"github.com/cmj0121/logger"
)

func main() {
	log := logger.New("example")
	log.SetLevel("info")

	log.Crit("show the CRIT log")
	log.Warn("show the WARN log")
	log.Info("show the INFO log")
	log.Debug("show the DEBUG log")
	log.Verbose("show the VERBOSE log")
}
