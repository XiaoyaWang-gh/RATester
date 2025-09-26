package hugofs

import (
	"fmt"
	"testing"
	"time"
)

func TestModTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &dirNameOnlyFileInfo{
		name:    "test",
		modTime: time.Now(),
	}
	if got := fi.ModTime(); !got.Equal(fi.modTime) {
		t.Errorf("ModTime() = %v, want %v", got, fi.modTime)
	}
}
