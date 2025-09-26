package port

import (
	"fmt"
	"testing"
)

func TestString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	builder := &nameBuilder{
		name:          "test",
		rangePortFrom: 1000,
		rangePortTo:   1001,
	}
	if builder.String() != "Port_test_1000_1001" {
		t.Errorf("String() = %v, want %v", builder.String(), "Port_test_1000_1001")
	}
}
