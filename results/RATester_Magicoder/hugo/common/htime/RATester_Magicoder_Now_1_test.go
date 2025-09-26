package htime

import (
	"fmt"
	"testing"
	"time"
)

func TestNow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := Now()

	if now.IsZero() {
		t.Error("Now() should not return a zero time")
	}

	if now.Location() != time.UTC {
		t.Error("Now() should return a time in UTC")
	}
}
