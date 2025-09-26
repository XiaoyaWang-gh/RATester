package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	writer := &SLACKWriter{}
	logMsg := &LogMsg{
		When: time.Now(),
	}

	formatted := writer.Format(logMsg)

	if formatted == "" {
		t.Error("Expected a non-empty string, got an empty string")
	}
}
