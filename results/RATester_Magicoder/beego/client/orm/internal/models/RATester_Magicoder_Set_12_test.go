package models

import (
	"fmt"
	"testing"
	"time"
)

func TestSet_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := time.Now()
	df := DateField{}
	df.Set(now)

	if df.Value() != now {
		t.Errorf("Expected %v, got %v", now, df.Value())
	}
}
