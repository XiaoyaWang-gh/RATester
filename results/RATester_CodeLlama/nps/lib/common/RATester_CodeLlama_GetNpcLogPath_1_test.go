package common

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestGetNpcLogPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if IsWindows() {
		if GetNpcLogPath() != filepath.Join(GetAppPath(), "npc.log") {
			t.Errorf("GetNpcLogPath() = %v, want %v", GetNpcLogPath(), filepath.Join(GetAppPath(), "npc.log"))
		}
	} else {
		if GetNpcLogPath() != "/var/log/npc.log" {
			t.Errorf("GetNpcLogPath() = %v, want %v", GetNpcLogPath(), "/var/log/npc.log")
		}
	}
}
