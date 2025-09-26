package logs

import (
	"fmt"
	"testing"
)

func TestResetColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if ResetColor() != "\033[0m" {
		t.Errorf("ResetColor() = %q, want %q", ResetColor(), "\033[0m")
	}
}
