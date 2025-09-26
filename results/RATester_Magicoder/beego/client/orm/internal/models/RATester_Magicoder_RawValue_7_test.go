package models

import (
	"fmt"
	"testing"
	"time"
)

func TestRawValue_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := time.Now()
	field := DateTimeField(now)

	rawValue := field.RawValue()

	if rawValue != now {
		t.Errorf("Expected %v, got %v", now, rawValue)
	}
}
