package daemon

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestStart_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	osArgs := []string{"test"}
	f := "test"
	pidPath := "/tmp"
	runPath := "/tmp"

	start(osArgs, f, pidPath, runPath)

	pidFile := filepath.Join(pidPath, f+".pid")
	_, err := os.Stat(pidFile)
	if err != nil {
		t.Errorf("Expected pid file to exist, but it does not")
	}

	pid, err := ioutil.ReadFile(pidFile)
	if err != nil {
		t.Errorf("Expected to be able to read pid file, but could not")
	}

	cmd := exec.Command("kill", "-0", string(pid))
	err = cmd.Run()
	if err != nil {
		t.Errorf("Expected process to be running, but it is not")
	}

	os.Remove(pidFile)
}
