package common

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestGetLogPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if IsWindows() {
		if GetLogPath() != filepath.Join(GetAppPath(), "nps.log") {
			t.Errorf("GetLogPath() = %v, want %v", GetLogPath(), filepath.Join(GetAppPath(), "nps.log"))
		}
	} else {
		if GetLogPath() != "/var/log/nps.log" {
			t.Errorf("GetLogPath() = %v, want %v", GetLogPath(), "/var/log/nps.log")
		}
	}
}
