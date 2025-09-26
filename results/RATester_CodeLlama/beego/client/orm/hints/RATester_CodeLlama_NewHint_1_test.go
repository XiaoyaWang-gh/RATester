package hints

import (
	"fmt"
	"testing"
)

func TestNewHint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "key"
	value := "value"
	hint := NewHint(key, value)
	if hint.GetKey() != key {
		t.Errorf("NewHint() = %v, want %v", hint.GetKey(), key)
	}
	if hint.GetValue() != value {
		t.Errorf("NewHint() = %v, want %v", hint.GetValue(), value)
	}
}
