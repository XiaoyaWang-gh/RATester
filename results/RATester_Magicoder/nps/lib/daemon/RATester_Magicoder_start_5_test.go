package daemon

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
)

func Teststart_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	osArgs := []string{"test"}
	f := "test"
	pidPath := "/tmp/test"
	runPath := "/tmp/test"

	start(osArgs, f, pidPath, runPath)

	// Test if the function correctly writes the PID to the file
	pidFile := filepath.Join(pidPath, f+".pid")
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		t.Errorf("Error reading PID file: %v", err)
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		t.Errorf("Error converting PID to int: %v", err)
	}
	if pid <= 0 {
		t.Errorf("Invalid PID: %d", pid)
	}

	// Test if the function correctly starts the command
	cmd := exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "pid=")
	output, err := cmd.Output()
	if err != nil {
		t.Errorf("Error checking if process is running: %v", err)
	}
	if string(output) != strconv.Itoa(pid)+"\n" {
		t.Errorf("Process not running")
	}

	// Clean up
	cmd = exec.Command("kill", strconv.Itoa(pid))
	cmd.Run()
	os.Remove(pidFile)
}
