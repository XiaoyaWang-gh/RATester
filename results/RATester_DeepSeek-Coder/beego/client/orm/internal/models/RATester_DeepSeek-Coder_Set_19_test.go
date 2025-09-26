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

	testTime := time.Now()
	e := &TimeField{}
	e.Set(testTime)

	if *e != TimeField(testTime) {
		t.Errorf("Expected %v, got %v", testTime, *e)
	}
}
