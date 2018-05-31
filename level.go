package yolo

// Level is the log level
type Level int

const (
	// Log levels
	LevelDebug Level = iota + 1
	LevelInfo
	LevelWarn
	LevelError
)

// String returns the log level in string representation
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// LevelFromString returns the log level from string representation
func LevelFromString(s string) (Level, error) {
	switch s {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	default:
		return 0, ErrUnknownLevel
	}
}
