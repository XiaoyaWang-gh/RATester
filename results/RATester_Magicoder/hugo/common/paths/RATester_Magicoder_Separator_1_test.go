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

	bridge := pathBridge{}
	separator := bridge.Separator()
	if separator != "/" {
		t.Errorf("Expected separator to be '/', but got %s", separator)
	}
}
