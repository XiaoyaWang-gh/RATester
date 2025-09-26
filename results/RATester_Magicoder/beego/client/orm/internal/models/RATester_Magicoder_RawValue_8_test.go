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

	now := time.Now()
	field := TimeField(now)

	rawValue := field.RawValue()

	if rawValue != now {
		t.Errorf("Expected RawValue to return %v, but got %v", now, rawValue)
	}
}
