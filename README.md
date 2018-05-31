![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)
![Tag](https://img.shields.io/github/tag/disq/yolo.svg)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/disq/yolo)
[![Go Report](https://goreportcard.com/badge/github.com/disq/yolo)](https://goreportcard.com/report/github.com/disq/yolo)

# yolo

Yet another logging library.

# Log Levels

	LevelDebug
	LevelInfo
	LevelWarning
	LevelError

# Options

The constructor accepts several options:

    WithLevel(level Level)                     // Set level
    WithOutput(output io.Writer)               // Set output
    WithPrefix(prefix string)                  // Set prefix
    WithTimeFormat(format string, utc bool)    // Set time format and UTC flag
    WithCaller(enabled bool)                   // Set call stack flag with file names
    WithCallerLong(enabled bool)               // Set call stack flag with file paths

## Defaults

If no options are given, the following are assumed.

    WithLevel(LevelInfo)
    WithOutput(os.Stdout)
    WithTimeFormat(DefaultTimeFormat, true)
    WithCaller(true)

The default time format is `2006-01-02 15:04:05.000`.

# Usage

```go
log := yolo.New() // Use defaults
log.Infof("Info message")
// 2017-12-21 22:23:24.256 INFO Info message

log = yolo.New(yolo.WithPrefix("[some-prefix]")) // constructor with optional prefix
log.Infof("Info message")
// 2017-12-21 22:23:24.256 INFO [some-prefix] Info message

// Create sub-logger, appending prefix
sublog := yolo.NewFrom(l, "[more-prefix]")
sublog.Errorf("Error message: %v", err)
//  2017-12-21 23:24:25.267 ERROR [some-prefix] [more-prefix] Error message: No such file or directory
```

# Helpers

Use `yolo.LevelFromString` to parse a string into a log level. This can be used like:

```go
// ...
l := flag.String("logLevel", "debug", "Log level")
flag.Parse()

lvl, err := yolo.LevelFromString(*l)
if err != nil {
	// Unknown log level
}

logger := yolo.New(yolo.WithLevel(lvl))
logger.Infof("One two three")
// ...
```