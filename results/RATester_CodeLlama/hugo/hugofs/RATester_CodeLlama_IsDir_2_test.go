package hugofs

import (
	"fmt"
	"testing"
)

func TestIsDir_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &dirNameOnlyFileInfo{
		name: "test",
	}
	if !fi.IsDir() {
		t.Errorf("IsDir() = %v, want %v", fi.IsDir(), true)
	}
}
