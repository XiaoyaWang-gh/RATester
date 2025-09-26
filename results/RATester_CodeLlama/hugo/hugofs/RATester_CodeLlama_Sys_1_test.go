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
	if got := fi.Sys(); got != nil {
		t.Errorf("Sys() = %v, want nil", got)
	}
}
