package models

import (
	"fmt"
	"testing"
	"time"
)

func TestSet_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var e *TimeField
	var d time.Time
	e.Set(d)
	if e.FieldType() != 0 {
		t.Errorf("FieldType() = %v, want %v", e.FieldType(), 0)
	}
	if e.RawValue() != nil {
		t.Errorf("RawValue() = %v, want %v", e.RawValue(), nil)
	}
	if e.SetRaw(d) != nil {
		t.Errorf("SetRaw() = %v, want %v", e.SetRaw(d), nil)
	}
	if e.String() != "" {
		t.Errorf("String() = %v, want %v", e.String(), "")
	}
	if e.Value() != d {
		t.Errorf("Value() = %v, want %v", e.Value(), d)
	}
}
