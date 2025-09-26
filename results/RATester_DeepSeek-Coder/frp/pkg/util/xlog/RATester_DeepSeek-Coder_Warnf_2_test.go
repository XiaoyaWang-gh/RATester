package xlog

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestWarnf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &Logger{
		prefixes: []LogPrefix{
			{Name: "test", Value: "test_value"},
		},
	}
	logger.renderPrefixString()
	expected := "[test=test_value] "

	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	logger.Warnf("test message")

	got := buf.String()
	if !strings.HasPrefix(got, expected) {
		t.Errorf("expected log message to start with %q, got %q", expected, got)
	}
}
