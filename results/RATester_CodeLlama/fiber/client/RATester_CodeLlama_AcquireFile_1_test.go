package client

import (
	"fmt"
	"testing"
)

func TestAcquireFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := AcquireFile()
	defer ReleaseFile(f)
	if f == nil {
		t.Error("AcquireFile() returns nil")
	}
}
