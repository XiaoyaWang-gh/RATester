package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TEST CASE CONTEXT_0
	b := &ValueBinder{}
	// TEST CASE CONTEXT_1
	// TEST CASE CONTEXT_2
	sourceParam := "sourceParam"
	// TEST CASE CONTEXT_3
	dest := &time.Time{}
	// TEST CASE CONTEXT_4
	layout := "layout"
	// TEST CASE CONTEXT_5
	valueMustExist := true
	// TEST CASE CONTEXT_6
	// TEST CASE CONTEXT_7
	b.time(sourceParam, dest, layout, valueMustExist)
	// TEST CASE CONTEXT_8
	if b.errors != nil {
		t.Errorf("errors should be nil")
	}
	// TEST CASE CONTEXT_9
	if dest == nil {
		t.Errorf("dest should not be nil")
	}
	// TEST CASE CONTEXT_10
	if dest.IsZero() {
		t.Errorf("dest should not be zero")
	}
}
