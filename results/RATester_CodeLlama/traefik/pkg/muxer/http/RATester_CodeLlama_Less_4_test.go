package http

import (
	"fmt"
	"testing"
)

func TestLess_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := routes{
		{priority: 1},
		{priority: 2},
		{priority: 3},
	}
	if !r.Less(0, 1) {
		t.Errorf("expected true")
	}
	if r.Less(1, 0) {
		t.Errorf("expected false")
	}
	if r.Less(0, 2) {
		t.Errorf("expected false")
	}
	if !r.Less(1, 2) {
		t.Errorf("expected true")
	}
	if r.Less(2, 1) {
		t.Errorf("expected false")
	}
	if r.Less(2, 0) {
		t.Errorf("expected false")
	}
	if !r.Less(1, 0) {
		t.Errorf("expected true")
	}
}
