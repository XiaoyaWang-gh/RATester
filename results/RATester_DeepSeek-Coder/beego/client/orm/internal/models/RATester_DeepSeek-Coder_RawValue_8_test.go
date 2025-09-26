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
	e := &TimeField{}
	e.Set(now)
	expected := now
	actual := e.RawValue()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
