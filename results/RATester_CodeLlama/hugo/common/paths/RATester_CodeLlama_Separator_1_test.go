package paths

import (
	"fmt"
	"testing"
)

func TestSeparator_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pathBridge := pathBridge{}
	if pathBridge.Separator() != "/" {
		t.Errorf("Separator() = %v, want %v", pathBridge.Separator(), "/")
	}
}
