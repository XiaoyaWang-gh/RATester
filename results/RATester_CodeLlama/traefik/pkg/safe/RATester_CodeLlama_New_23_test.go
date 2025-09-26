package safe

import (
	"fmt"
	"testing"
)

func TestNew_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	value := "hello"
	safe := New(value)
	if safe.value != value {
		t.Errorf("safe.value = %v, want %v", safe.value, value)
	}
}
