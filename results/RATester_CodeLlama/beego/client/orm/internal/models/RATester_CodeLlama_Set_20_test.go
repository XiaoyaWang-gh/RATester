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

	var e *DateTimeField
	d := time.Now()
	e.Set(d)
	if e.Value().String() != d.String() {
		t.Errorf("e.Value().String() = %v, want %v", e.Value().String(), d.String())
	}
}
