package schema

import (
	"fmt"
	"testing"
)

func TestNewEncoder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	encoder := NewEncoder()
	if encoder == nil {
		t.Error("NewEncoder() returned nil")
	}
	if encoder.cache == nil {
		t.Error("NewEncoder() did not initialize cache")
	}
	if len(encoder.regenc) != 0 {
		t.Error("NewEncoder() did not initialize regenc correctly")
	}
}
