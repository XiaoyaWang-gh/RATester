package gin

import (
	"fmt"
	"testing"
)

func TestForceConsoleColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ForceConsoleColor()
	if consoleColorMode != forceColor {
		t.Errorf("Expected consoleColorMode to be %v, but got %v", forceColor, consoleColorMode)
	}
}
