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

	fi := &dirNameOnlyFileInfo{
		name: "test",
	}
	if got := fi.Mode(); got != os.ModeDir {
		t.Errorf("dirNameOnlyFileInfo.Mode() = %v, want %v", got, os.ModeDir)
	}
}
