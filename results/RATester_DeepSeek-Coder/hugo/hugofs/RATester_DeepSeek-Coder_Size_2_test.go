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

	size := fi.Size()
	if size != 0 {
		t.Errorf("Expected size to be 0, got %d", size)
	}
}
