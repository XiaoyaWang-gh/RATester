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

	e := TimeField(time.Now())
	if e.Value().IsZero() {
		t.Errorf("TimeField.Value() = %v, want not zero", e.Value())
	}
}
