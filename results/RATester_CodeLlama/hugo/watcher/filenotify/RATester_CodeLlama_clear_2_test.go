package filenotify

import (
	"fmt"
	"os"
	"testing"
)

func TestClear_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &recording{}
	r.FileInfo = nil
	r.entries = map[string]os.FileInfo{}
	r.clear()
	if r.FileInfo != nil {
		t.Errorf("r.FileInfo = %v, want nil", r.FileInfo)
	}
	if r.entries != nil {
		t.Errorf("r.entries = %v, want nil", r.entries)
	}
}
