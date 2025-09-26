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

	now := time.Now()
	e := TimeField{}
	e.Set(now)

	if e.Value() != now {
		t.Errorf("Expected %v, got %v", now, e.Value())
	}
}
