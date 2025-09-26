package hugofs

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &dirNameOnlyFileInfo{
		name:    "test",
		modTime: time.Now(),
	}

	expectedMode := os.ModeDir
	actualMode := fi.Mode()

	if actualMode != expectedMode {
		t.Errorf("Expected mode %v, but got %v", expectedMode, actualMode)
	}
}
