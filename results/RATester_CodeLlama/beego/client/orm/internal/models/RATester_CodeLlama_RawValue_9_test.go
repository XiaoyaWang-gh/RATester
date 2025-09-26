package models

import (
	"fmt"
	"testing"
)

func TestRawValue_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	j := JSONField("{\"name\":\"john\"}")
	if j.RawValue() != "{\"name\":\"john\"}" {
		t.Errorf("RawValue() = %v, want %v", j.RawValue(), "{\"name\":\"john\"}")
	}
}
