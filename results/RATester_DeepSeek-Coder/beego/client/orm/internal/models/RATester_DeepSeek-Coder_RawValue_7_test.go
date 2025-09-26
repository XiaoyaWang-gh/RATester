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
	dt := DateTimeField(now)

	expected := now
	actual := dt.RawValue()

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
