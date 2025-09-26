package models

import (
	"fmt"
	"testing"
	"time"
)

func TestSet_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTime := time.Now()
	dt := DateTimeField{}
	dt.Set(testTime)

	if dt.Value() != testTime {
		t.Errorf("Expected %v, got %v", testTime, dt.Value())
	}
}
