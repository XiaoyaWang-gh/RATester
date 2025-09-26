package session

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
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
		savePath:    "/tmp/sessions",
	}

	sid := "test_session_id"
	ctx := context.Background()

	// Create a test file
	filePath := filepath.Join(fp.savePath, string(sid[0]), string(sid[1]), sid)
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test SessionDestroy
	err = fp.SessionDestroy(ctx, sid)
	if err != nil {
		t.Errorf("SessionDestroy returned an error: %v", err)
	}

	// Check if the file was deleted
	_, err = os.Stat(filePath)
	if err == nil {
		t.Errorf("SessionDestroy did not delete the file")
	}
}
