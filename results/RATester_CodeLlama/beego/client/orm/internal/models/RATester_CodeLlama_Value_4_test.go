package models

import (
	"fmt"
	"testing"
)

func TestValue_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	j := JsonbField("{\"name\":\"test\"}")
	if j.Value() != "{\"name\":\"test\"}" {
		t.Errorf("JsonbField.Value() = %v, want %v", j.Value(), "{\"name\":\"test\"}")
	}
}
