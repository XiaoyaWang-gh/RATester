package fiber

import (
	"fmt"
	"testing"
)

func TestWith_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	key := "key"
	value := "value"
	level := []uint8{1}
	r := &Redirect{}

	// When
	r.With(key, value, level...)

	// Then
	if len(r.messages) != 1 {
		t.Errorf("Expected %v, got %v", 1, len(r.messages))
	}
	if r.messages[0].key != key {
		t.Errorf("Expected %v, got %v", key, r.messages[0].key)
	}
	if r.messages[0].value != value {
		t.Errorf("Expected %v, got %v", value, r.messages[0].value)
	}
	if r.messages[0].level != level[0] {
		t.Errorf("Expected %v, got %v", level[0], r.messages[0].level)
	}
}
