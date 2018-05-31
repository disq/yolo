package yolo

import "fmt"

// Log is a generic logger interface
type Logger interface {
	Debug(i ...interface{})
	Debugf(format string, args ...interface{})
	Info(i ...interface{})
	Infof(format string, args ...interface{})
	Warn(i ...interface{})
	Warnf(format string, args ...interface{})
	Error(i ...interface{})
	Errorf(format string, args ...interface{})
}

// DefaultTimeFormat is our default time format
const DefaultTimeFormat = "2006-01-02 15:04:05.000"

// ErrUnknownLevel is returned by LevelFromString
var ErrUnknownLevel = fmt.Errorf("Unknown log level")
