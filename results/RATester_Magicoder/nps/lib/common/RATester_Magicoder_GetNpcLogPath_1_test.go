package common

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetNpcLogPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Setenv("OS", "windows")
	defer os.Unsetenv("OS")

	expected := filepath.Join(GetAppPath(), "npc.log")
	actual := GetNpcLogPath()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}

	os.Setenv("OS", "linux")
	expected = "/var/log/npc.log"
	actual = GetNpcLogPath()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
