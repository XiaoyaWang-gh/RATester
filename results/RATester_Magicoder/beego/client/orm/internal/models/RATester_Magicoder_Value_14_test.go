package models

import (
	"fmt"
	"testing"
	"time"
)

func TestValue_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := time.Now()
	dt := DateTimeField(now)

	if dt.Value() != now {
		t.Errorf("Expected %v, got %v", now, dt.Value())
	}
}
