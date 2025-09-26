package xlog

import (
	"fmt"
	"testing"
)

func TestAppendPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &Logger{}
	prefix := "test_prefix"
	logger.AppendPrefix(prefix)

	if len(logger.prefixes) != 1 {
		t.Errorf("Expected 1 prefix, got %d", len(logger.prefixes))
	}

	if logger.prefixes[0].Name != prefix || logger.prefixes[0].Value != prefix {
		t.Errorf("Expected prefix name and value to be %s, got %s", prefix, logger.prefixes[0].Name)
	}

	if logger.prefixString != prefix {
		t.Errorf("Expected prefix string to be %s, got %s", prefix, logger.prefixString)
	}
}
