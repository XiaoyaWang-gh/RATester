package models

import (
	"fmt"
	"testing"
)

func TestValue_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := IntegerField(-2147483648)
	if e.Value() != -2147483648 {
		t.Errorf("e.Value() = %v, want %v", e.Value(), -2147483648)
	}
}
