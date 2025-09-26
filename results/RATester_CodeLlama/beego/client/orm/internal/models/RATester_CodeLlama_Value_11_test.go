package models

import (
	"fmt"
	"testing"
)

func TestValue_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := BigIntegerField(-9223372036854775808)
	if e.Value() != -9223372036854775808 {
		t.Errorf("e.Value() = %v, want %v", e.Value(), -9223372036854775808)
	}
}
