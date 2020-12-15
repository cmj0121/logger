package logger

// the program info
const (
	ENV_LOG_LEVEL = "LOG_LEVEL"
	PROJ_NAME     = "Logger"

	MAJOR = 1
	MINOR = 2
	MACRO = 1
)

const (
	CRIT LogLevel = iota
	WARN
	INFO
	DEBUG
	VERBOSE
)

var (
	// the human-readable log level
	log_level = map[string]LogLevel{
		"CRIT":    CRIT,
		"WARN":    WARN,
		"INFO":    INFO,
		"DEBUG":   DEBUG,
		"VERBOSE": VERBOSE,
	}
)
