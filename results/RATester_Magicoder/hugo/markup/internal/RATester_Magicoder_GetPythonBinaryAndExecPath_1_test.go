package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetPythonBinaryAndExecPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pythonBinary, pythonExecPath := GetPythonBinaryAndExecPath()

	if pythonBinary == "" || pythonExecPath == "" {
		t.Error("Python binary and exec path not found")
	}

	if _, err := os.Stat(pythonExecPath); os.IsNotExist(err) {
		t.Error("Python exec path does not exist")
	}

	if filepath.Base(pythonBinary) != pythonExecPath {
		t.Error("Python binary and exec path do not match")
	}
}
