package yolo

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Yolo is our logger struct
type Yolo struct {
	mu   sync.Mutex // used to synchronize output as well
	opts *options
}

// New creates a new Yolo
func New(opts ...Option) *Yolo {
	loggerOpts := options{
		level:      LevelInfo,
		output:     os.Stdout,
		timeFormat: DefaultTimeFormat,
		timeUTC:    true,
		caller:     true,
	}

	if len(opts) != 0 {
		for _, opt := range opts {
			opt(&loggerOpts)
		}
	}

	return &Yolo{
		opts: &loggerOpts,
	}
}

// NewFrom creates a new logger from a logger, appending the prefix
func NewFrom(from *Yolo, morePrefix string) *Yolo {
	pf := from.opts.prefix
	if morePrefix != "" {
		pf += morePrefix + " "
	}

	loggerOpts := options{
		prefix: pf,

		level:      from.opts.level,
		output:     from.opts.output,
		caller:     from.opts.caller,
		callerLong: from.opts.callerLong,
		timeFormat: from.opts.timeFormat,
		timeUTC:    from.opts.timeUTC,
	}

	return &Yolo{
		opts: &loggerOpts,
	}
}

func (l *Yolo) fmt(level Level, msg string) string {
	line := strings.Builder{}

	if l.opts.timeFormat != "" {
		tm := time.Now()
		if l.opts.timeUTC {
			tm = tm.UTC()
		}

		line.WriteString(tm.Format(l.opts.timeFormat))
		line.WriteString(" ")
	}

	line.WriteString(level.String())
	line.WriteString(" ")
	line.WriteString(l.opts.prefix)

	if l.opts.caller {
		var ok bool
		_, file, cline, ok := runtime.Caller(3)
		if !ok {
			file = "???"
			cline = 0
		}

		if !l.opts.callerLong {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}

		line.WriteString(file)
		line.WriteString(":")
		line.WriteString(strconv.Itoa(cline))
		line.WriteString(": ")
	}

	line.WriteString(msg)
	line.WriteString("\n")

	return line.String()
}

func (l *Yolo) write(level Level, format string, a ...interface{}) {
	if l.opts.level > level || l.opts.output == nil {
		return
	}

	formatted := l.fmt(level, fmt.Sprintf(format, a...))

	l.mu.Lock()
	// ignore write errors
	_, _ = l.opts.output.Write([]byte(formatted))
	l.mu.Unlock()
}

// Debugf logs formatted message in debug level
func (l *Yolo) Debugf(format string, a ...interface{}) {
	l.write(LevelDebug, format, a...)
}

// Infof logs formatted message in info level
func (l *Yolo) Infof(format string, a ...interface{}) {
	l.write(LevelInfo, format, a...)
}

// Warnf logs formatted message in warn level
func (l *Yolo) Warnf(format string, a ...interface{}) {
	l.write(LevelWarn, format, a...)
}

// Errorf logs formatted message in error level
func (l *Yolo) Errorf(format string, a ...interface{}) {
	l.write(LevelError, format, a...)
}
