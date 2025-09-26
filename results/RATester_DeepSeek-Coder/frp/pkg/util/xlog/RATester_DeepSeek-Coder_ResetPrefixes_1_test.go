package xlog

import (
	"fmt"
	"testing"
)

func TestResetPrefixes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &Logger{
		prefixes: []LogPrefix{
			{Name: "test1", Value: "value1"},
			{Name: "test2", Value: "value2"},
		},
		prefixString: "test1:value1,test2:value2",
	}

	oldPrefixes := logger.ResetPrefixes()

	if len(logger.prefixes) != 0 {
		t.Errorf("Expected prefixes to be empty, got %v", logger.prefixes)
	}

	if logger.prefixString != "" {
		t.Errorf("Expected prefixString to be empty, got %v", logger.prefixString)
	}

	if len(oldPrefixes) != 2 {
		t.Errorf("Expected oldPrefixes to have 2 elements, got %v", oldPrefixes)
	}

	if oldPrefixes[0].Name != "test1" || oldPrefixes[0].Value != "value1" {
		t.Errorf("Expected first oldPrefix to be test1:value1, got %v:%v", oldPrefixes[0].Name, oldPrefixes[0].Value)
	}

	if oldPrefixes[1].Name != "test2" || oldPrefixes[1].Value != "value2" {
		t.Errorf("Expected second oldPrefix to be test2:value2, got %v:%v", oldPrefixes[1].Name, oldPrefixes[1].Value)
	}
}
