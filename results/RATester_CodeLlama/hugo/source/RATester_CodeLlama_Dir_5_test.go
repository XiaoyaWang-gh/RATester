package source

import (
	"fmt"
	"testing"
)

func TestDir_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &File{}
	if got := fi.Dir(); got != "" {
		t.Errorf("Dir() = %v, want \"\"", got)
	}
}
