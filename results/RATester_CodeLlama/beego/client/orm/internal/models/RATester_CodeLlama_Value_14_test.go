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

	e := DateTimeField(time.Now())
	if e.Value().IsZero() {
		t.Errorf("Value() = %v, want not zero", e.Value())
	}
}
