package ginS

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadHTMLGlob_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pattern := "./templates/*"
	LoadHTMLGlob(pattern)

	_, err := os.Stat(pattern)
	if os.IsNotExist(err) {
		t.Errorf("LoadHTMLGlob(%s) failed, file does not exist", pattern)
	}
}
