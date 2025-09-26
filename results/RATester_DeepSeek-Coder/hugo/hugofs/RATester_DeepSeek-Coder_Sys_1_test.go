package hugofs

import (
	"fmt"
	"testing"
	"time"
)

func TestSys_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &dirNameOnlyFileInfo{
		name:    "test",
		modTime: time.Now(),
	}

	result := fi.Sys()

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
