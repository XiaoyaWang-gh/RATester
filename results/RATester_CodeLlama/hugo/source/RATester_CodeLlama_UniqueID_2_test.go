package source

import (
	"fmt"
	"testing"
)

func TestUniqueID_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &File{}
	if got := fi.UniqueID(); got != "" {
		t.Errorf("UniqueID() = %v, want \"\"", got)
	}
}
