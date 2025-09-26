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
		t.Errorf("NewEncoder() = %v, want not nil", encoder)
	}
	if encoder.cache == nil {
		t.Errorf("NewEncoder().cache = %v, want not nil", encoder.cache)
	}
	if encoder.regenc == nil {
		t.Errorf("NewEncoder().regenc = %v, want not nil", encoder.regenc)
	}
}
