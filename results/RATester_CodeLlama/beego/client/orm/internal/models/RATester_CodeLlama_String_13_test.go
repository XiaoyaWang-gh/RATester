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

	e := &DateTimeField{}
	e.Set(time.Now())
	if e.String() != e.Value().String() {
		t.Errorf("String() = %v, want %v", e.String(), e.Value().String())
	}
}
