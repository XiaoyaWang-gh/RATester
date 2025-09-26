package gin

import (
	"fmt"
	"testing"
)

func TestDisableConsoleColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	DisableConsoleColor()
	if consoleColorMode != disableColor {
		t.Errorf("Expected consoleColorMode to be %v, but got %v", disableColor, consoleColorMode)
	}
}
