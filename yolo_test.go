package yolo_test

import (
	"bytes"
	"testing"

	"github.com/disq/yolo"
)

func TestBasic(t *testing.T) {

	{
		var b bytes.Buffer

		l := yolo.New(yolo.WithLevel(yolo.LevelDebug), yolo.WithTimeFormat("", false), yolo.WithPrefix("[test-debug]"), yolo.WithOutput(&b))
		l.Debugf("Debug message")
		l.Errorf("Error message")

		result := b.Bytes()
		golden := []byte(`DEBUG [test-debug] yolo_test.go:16: Debug message
ERROR [test-debug] yolo_test.go:17: Error message
`)
		if !bytes.Equal(golden, result) {
			t.Errorf("Basic test: Got: %q Want: %q", string(result), string(golden))
		}

		{
			golden = append(golden, []byte("DEBUG [test-debug] [a sublogger] yolo_test.go:29: Debug message\n")...)
			yolo.NewFrom(l, "[a sublogger]").Debugf("Debug message")
			result := b.Bytes()
			if !bytes.Equal(golden, result) {
				t.Errorf("Basic test: Got: %q Want: %q", string(result), string(golden))
			}
		}
	}
}

func TestLevel(t *testing.T) {
	var b bytes.Buffer

	l := yolo.New(yolo.WithLevel(yolo.LevelWarning), yolo.WithTimeFormat("", false), yolo.WithOutput(&b))
	l.Debugf("Debug message")
	l.Infof("Info message")
	l.Warningf("Warning message")
	l.Errorf("Error message")

	result := b.Bytes()
	golden := []byte("WARNING yolo_test.go:44: Warning message\nERROR yolo_test.go:45: Error message\n")
	if !bytes.Equal(golden, result) {
		t.Errorf("Got: %q Want: %q", string(result), string(golden))
	}
}

func TestCallerDisabled(t *testing.T) {
	var b bytes.Buffer

	l := yolo.New(yolo.WithCaller(false), yolo.WithTimeFormat("", false), yolo.WithOutput(&b))
	l.Infof("Info message")

	result := b.Bytes()
	golden := []byte("INFO Info message\n")
	if !bytes.Equal(golden, result) {
		t.Errorf("Got: %q Want: %q", string(result), string(golden))
	}
}
