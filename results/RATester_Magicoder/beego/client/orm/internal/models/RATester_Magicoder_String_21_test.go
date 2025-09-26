package models

import (
	"fmt"
	"testing"
	"time"
)

func TestString_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := time.Now()
	df := DateField(now)
	expected := now.Format("2006-01-02")
	result := df.String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
