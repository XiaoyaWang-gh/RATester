package xlog

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestDebugf_2(t *testing.T) {
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

	logger.Debugf("test message")

	output := buf.String()
	if !strings.HasPrefix(output, expected) {
		t.Errorf("Expected output to start with %s, got %s", expected, output)
	}
}
