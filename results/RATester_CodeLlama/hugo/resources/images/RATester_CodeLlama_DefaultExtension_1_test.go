package images

import (
	"fmt"
	"testing"
)

func TestDefaultExtension_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := Format(0)
	if got := f.DefaultExtension(); got != "" {
		t.Errorf("DefaultExtension() = %v, want \"\"", got)
	}
}
