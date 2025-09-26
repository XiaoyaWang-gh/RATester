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

	now := time.Now()
	df := DateField(now)

	if df.Value() != now {
		t.Errorf("Expected %v, got %v", now, df.Value())
	}
}
