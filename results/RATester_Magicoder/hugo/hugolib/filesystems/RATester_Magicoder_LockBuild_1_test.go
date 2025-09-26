package filesystems

import (
	"fmt"
	"testing"
)

func TestLockBuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BaseFs{}
	unlock, err := b.LockBuild()
	if err != nil {
		t.Fatalf("LockBuild() returned an error: %v", err)
	}
	defer unlock()

	// Add more tests here if needed.
}
