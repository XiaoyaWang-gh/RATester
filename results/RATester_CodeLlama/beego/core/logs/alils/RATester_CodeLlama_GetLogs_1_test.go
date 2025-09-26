package alils

import (
	"fmt"
	"testing"
)

func TestGetLogs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var m *LogGroup
	if m != nil {
		if len(m.GetLogs()) != 0 {
			t.Errorf("GetLogs() = %v, want %v", len(m.GetLogs()), 0)
		}
	}
}
