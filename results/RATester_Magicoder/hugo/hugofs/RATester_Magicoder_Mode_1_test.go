package hugofs

import (
	"fmt"
	"os"
	"testing"
)

func TestMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &dirNameOnlyFileInfo{name: "test"}
	mode := fi.Mode()
	if mode != os.ModeDir {
		t.Errorf("Expected ModeDir, got %v", mode)
	}
}
