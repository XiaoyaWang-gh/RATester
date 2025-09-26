package client

import (
	"fmt"
	"testing"
)

func TestReleaseFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &File{}
	ReleaseFile(f)
	if f.reader != nil {
		t.Errorf("ReleaseFile() reader should be nil")
	}
	if f.name != "" {
		t.Errorf("ReleaseFile() name should be empty")
	}
	if f.fieldName != "" {
		t.Errorf("ReleaseFile() fieldName should be empty")
	}
	if f.path != "" {
		t.Errorf("ReleaseFile() path should be empty")
	}
}
