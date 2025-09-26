package models

import (
	"fmt"
	"testing"
	"time"
)

func TestString_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := time.Now()
	dt := DateTimeField(now)

	expected := now.Format(time.RFC3339)
	actual := dt.String()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
