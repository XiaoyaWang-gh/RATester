package hints

import (
	"fmt"
	"testing"
)

func TestLimit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := int64(10)
	hint := Limit(d)

	if hint.GetKey() != "LIMIT" {
		t.Errorf("Expected 'LIMIT', got '%v'", hint.GetKey())
	}

	if hint.GetValue() != d {
		t.Errorf("Expected '%v', got '%v'", d, hint.GetValue())
	}
}
