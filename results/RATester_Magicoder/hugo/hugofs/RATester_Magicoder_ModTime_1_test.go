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

	testTime := time.Now()
	fi := &dirNameOnlyFileInfo{
		name:    "test",
		modTime: testTime,
	}

	if fi.ModTime() != testTime {
		t.Errorf("Expected ModTime to be %v, but got %v", testTime, fi.ModTime())
	}
}
