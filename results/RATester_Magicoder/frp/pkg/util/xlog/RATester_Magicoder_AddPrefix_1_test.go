package xlog

import (
	"fmt"
	"testing"
)

func TestAddPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &Logger{}
	prefix := LogPrefix{
		Name:     "test",
		Value:    "test_value",
		Priority: 1,
	}
	logger.AddPrefix(prefix)

	if len(logger.prefixes) != 1 {
		t.Errorf("Expected 1 prefix, got %d", len(logger.prefixes))
	}

	if logger.prefixes[0] != prefix {
		t.Errorf("Expected prefix %v, got %v", prefix, logger.prefixes[0])
	}

	if logger.prefixString != "test=test_value" {
		t.Errorf("Expected prefix string 'test=test_value', got '%s'", logger.prefixString)
	}
}
