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
		t.Error("New() should not return nil")
	}
	if len(logger.prefixes) != 0 {
		t.Error("New() should initialize prefixes as an empty slice")
	}
}
