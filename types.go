package yolo

import "fmt"

// Log is a generic logger interface
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

var _ Logger = (*Yolo)(nil)

// DefaultTimeFormat is our default time format
const DefaultTimeFormat = "2006-01-02 15:04:05.000"

// ErrUnknownLevel is returned by LevelFromString
var ErrUnknownLevel = fmt.Errorf("Unknown log level")
