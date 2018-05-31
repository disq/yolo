package yolo

import "io"

type options struct {
	level      Level
	output     io.Writer
	prefix     string
	timeFormat string
	timeUTC    bool
	caller     bool
	callerLong bool
}

// Option is our option type
type Option func(*options)

// WithLevel sets the log level
func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

// WithOutput sets log output
func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

// WithPrefix sets a log prefix
func WithPrefix(prefix string) Option {
	return func(o *options) {
		if prefix != "" {
			prefix += " "
		}
		o.prefix = prefix
	}
}

// WithCaller enables caller logging
func WithCaller(enabled bool) Option {
	return func(o *options) {
		o.caller = enabled
		o.callerLong = false
	}
}

// WithCallerLong enables caller logging, long format
func WithCallerLong(enabled bool) Option {
	return func(o *options) {
		o.caller = enabled
		o.callerLong = true
	}
}

// WithTimeFormat sets the time format. Specify empty time format to disable datetime in logs.
func WithTimeFormat(format string, utc bool) Option {
	return func(o *options) {
		o.timeFormat = format
		o.timeUTC = utc
	}
}
