package models

import (
	"fmt"
	"testing"
)

func TestSetRaw_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var j JSONField
	var value interface{}
	var err error

	value = "hello"
	err = j.SetRaw(value)
	if err != nil {
		t.Errorf("SetRaw failed: %v", err)
	}
	if j.RawValue() != value {
		t.Errorf("SetRaw failed: %v", j.RawValue())
	}

	value = 123
	err = j.SetRaw(value)
	if err == nil {
		t.Errorf("SetRaw failed: %v", err)
	}
}
