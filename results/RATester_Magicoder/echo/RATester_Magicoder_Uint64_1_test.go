package echo

import (
	"fmt"
	"testing"
)

func TestUint64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			// Implement your test case here
			return "123"
		},
	}

	var dest uint64
	err := b.Uint64("test", &dest)

	if err != nil {
		t.Errorf("Uint64() error = %v, wantErr %v", err, false)
		return
	}

	if dest != 123 {
		t.Errorf("Uint64() = %v, want %v", dest, 123)
	}
}
