package session

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"testing"
)

func TestSessionDestroy_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fp := &FileProvider{
		lock:        sync.RWMutex{},
		maxlifetime: 100,
		savePath:    "/tmp",
	}

	sid := "test_session_id"

	// Create a temporary file
	file, err := ioutil.TempFile(fp.savePath, "")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write some data to the file
	_, err = file.WriteString("test data")
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	// Test SessionDestroy
	err = fp.SessionDestroy(context.Background(), sid)
	if err != nil {
		t.Errorf("SessionDestroy failed: %v", err)
	}

	// Check if the file was deleted
	_, err = os.Stat(file.Name())
	if !os.IsNotExist(err) {
		t.Errorf("File was not deleted: %v", err)
	}
}
