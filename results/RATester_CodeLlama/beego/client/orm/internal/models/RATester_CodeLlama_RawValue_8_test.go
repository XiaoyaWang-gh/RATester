package models

import (
	"fmt"
	"testing"
	"time"
)

func TestRawValue_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var e *TimeField
	e.Set(time.Now())
	if e.RawValue() != e.Value() {
		t.Errorf("RawValue() = %v, want %v", e.RawValue(), e.Value())
	}
}
