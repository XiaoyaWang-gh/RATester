package common

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetInstallPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if runtime.GOOS == "windows" {
		if GetInstallPath() != `C:\Program Files\nps` {
			t.Errorf("GetInstallPath() = %v, want %v", GetInstallPath(), `C:\Program Files\nps`)
		}
	} else {
		if GetInstallPath() != "/etc/nps" {
			t.Errorf("GetInstallPath() = %v, want %v", GetInstallPath(), "/etc/nps")
		}
	}
}
