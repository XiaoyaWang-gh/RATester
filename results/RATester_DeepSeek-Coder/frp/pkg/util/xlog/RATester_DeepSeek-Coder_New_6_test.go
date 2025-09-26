package xlog

import (
	"fmt"
	"testing"
)

func TestNew_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := New()
	if logger == nil {
		t.Errorf("New() = %v, want not nil", logger)
	}
	if len(logger.prefixes) != 0 {
		t.Errorf("New().prefixes = %v, want empty", logger.prefixes)
	}
}
