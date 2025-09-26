package htesting

import (
	"fmt"
	"testing"
)

func TestGoMinorVersion_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if GoMinorVersion() != 10 {
		t.Errorf("GoMinorVersion() = %d, want %d", GoMinorVersion(), 10)
	}
}
