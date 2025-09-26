package allconfig

import (
	"fmt"
	"testing"
)

func TestIgnoredLogs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Config{
		C: &ConfigCompiled{
			IgnoredLogs: map[string]bool{
				"test": true,
			},
		},
	}

	configLanguage := ConfigLanguage{
		config: config,
	}

	ignoredLogs := configLanguage.IgnoredLogs()

	if len(ignoredLogs) != 1 {
		t.Errorf("Expected 1 ignored log, but got %d", len(ignoredLogs))
	}

	if _, ok := ignoredLogs["test"]; !ok {
		t.Error("Expected to find 'test' in ignored logs")
	}
}
