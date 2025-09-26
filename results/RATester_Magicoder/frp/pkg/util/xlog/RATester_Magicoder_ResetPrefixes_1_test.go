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

	old := logger.ResetPrefixes()

	if len(logger.prefixes) != 0 {
		t.Errorf("Expected prefixes to be empty, but got %v", logger.prefixes)
	}

	if logger.prefixString != "" {
		t.Errorf("Expected prefixString to be empty, but got %s", logger.prefixString)
	}

	if len(old) != 2 {
		t.Errorf("Expected old prefixes to have length 2, but got %v", old)
	}

	if old[0].Name != "test1" || old[0].Value != "value1" {
		t.Errorf("Expected first old prefix to be test1:value1, but got %s:%s", old[0].Name, old[0].Value)
	}

	if old[1].Name != "test2" || old[1].Value != "value2" {
		t.Errorf("Expected second old prefix to be test2:value2, but got %s:%s", old[1].Name, old[1].Value)
	}
}
