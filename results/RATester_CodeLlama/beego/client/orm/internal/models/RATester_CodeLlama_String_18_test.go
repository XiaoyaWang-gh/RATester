package models

import (
	"fmt"
	"testing"
)

func TestString_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	j := JSONField("{\"name\":\"john\"}")
	if j.String() != "{\"name\":\"john\"}" {
		t.Errorf("JSONField.String() = %v, want %v", j.String(), "{\"name\":\"john\"}")
	}
}
