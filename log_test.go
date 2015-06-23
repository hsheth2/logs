package logs

import (
	"testing"
)

// TODO write better tests

func TestLoggers(t *testing.T) {
	Trace.Println("Test Trace")
	Info.Println("Test Info")
	Error.Println("Test Error")
}
