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

	t.Parallel()

	bridge := pathBridge{}
	expected := "/"
	result := bridge.Separator()

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
