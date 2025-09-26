package models

import (
	"fmt"
	"testing"
)

func TestValue_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	j := JSONField("{\"name\":\"test\"}")
	if j.Value() != "{\"name\":\"test\"}" {
		t.Errorf("JSONField.Value() = %v, want %v", j.Value(), "{\"name\":\"test\"}")
	}
}
