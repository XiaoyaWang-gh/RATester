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

	logger := &Logger{
		prefixes: []LogPrefix{
			{
				Name:     "test",
				Value:    "test_value",
				Priority: 10,
			},
		},
	}

	newPrefix := LogPrefix{
		Name:     "new_test",
		Value:    "new_test_value",
		Priority: 20,
	}

	logger.AddPrefix(newPrefix)

	if len(logger.prefixes) != 2 {
		t.Errorf("Expected length of prefixes to be 2, got %d", len(logger.prefixes))
	}

	if logger.prefixes[1].Name != "new_test" {
		t.Errorf("Expected second prefix name to be 'new_test', got %s", logger.prefixes[1].Name)
	}

	if logger.prefixes[1].Value != "new_test_value" {
		t.Errorf("Expected second prefix value to be 'new_test_value', got %s", logger.prefixes[1].Value)
	}

	if logger.prefixes[1].Priority != 20 {
		t.Errorf("Expected second prefix priority to be 20, got %d", logger.prefixes[1].Priority)
	}
}
