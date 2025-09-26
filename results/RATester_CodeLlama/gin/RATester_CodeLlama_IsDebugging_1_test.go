package gin

import (
	"fmt"
	"testing"
)

func TestIsDebugging_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ginMode = debugCode
	if !IsDebugging() {
		t.Errorf("IsDebugging() = false, want true")
	}
}
