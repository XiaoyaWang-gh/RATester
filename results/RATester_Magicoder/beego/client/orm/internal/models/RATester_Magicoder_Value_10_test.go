package models

import (
	"fmt"
	"testing"
	"time"
)

func TestValue_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := time.Now()
	field := TimeField(now)

	value := field.Value()

	if value != now {
		t.Errorf("Expected %v, got %v", now, value)
	}
}
