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
	fileInfo := &dirNameOnlyFileInfo{
		name:    "testFile",
		modTime: testTime,
	}

	if fileInfo.ModTime() != testTime {
		t.Errorf("Expected ModTime to be %v, got %v", testTime, fileInfo.ModTime())
	}
}
