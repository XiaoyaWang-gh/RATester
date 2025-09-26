package hugofs

import (
	"fmt"
	"testing"
)

func TestSize_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &dirNameOnlyFileInfo{
		name: "test",
	}
	if fi.Size() != 0 {
		t.Errorf("Size() = %v, want %v", fi.Size(), 0)
	}
}
