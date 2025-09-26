package logs

import (
	"fmt"
	"testing"
)

func TestFormat_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &JLWriter{}
	lm := &LogMsg{}
	msg := s.Format(lm)
	if msg != "" {
		t.Errorf("Expected empty string, got %s", msg)
	}
}
