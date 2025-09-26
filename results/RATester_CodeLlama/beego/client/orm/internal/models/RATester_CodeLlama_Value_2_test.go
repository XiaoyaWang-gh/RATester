package models

import (
	"fmt"
	"testing"
	"time"
)

func TestValue_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var e DateField
	if e.Value() != time.Time(e) {
		t.Errorf("e.Value() = %v, want %v", e.Value(), time.Time(e))
	}
}
