package source

import (
	"fmt"
	"testing"
)

func TestLogicalName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &File{}
	if got := fi.LogicalName(); got != "" {
		t.Errorf("LogicalName() = %v, want %v", got, "")
	}
}
