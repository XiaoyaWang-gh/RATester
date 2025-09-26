package web

import (
	"fmt"
	"testing"
)

func TestLe_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var arg1 int
	var arg2 int
	var expected bool
	var actual bool
	var err error

	arg1 = 1
	arg2 = 2
	expected = true
	actual, err = le(arg1, arg2)
	if err != nil {
		t.Errorf("le(%v, %v) error: %v", arg1, arg2, err)
		return
	}
	if actual != expected {
		t.Errorf("le(%v, %v) = %v, expected %v", arg1, arg2, actual, expected)
	}

	arg1 = 2
	arg2 = 2
	expected = true
	actual, err = le(arg1, arg2)
	if err != nil {
		t.Errorf("le(%v, %v) error: %v", arg1, arg2, err)
		return
	}
	if actual != expected {
		t.Errorf("le(%v, %v) = %v, expected %v", arg1, arg2, actual, expected)
	}

	arg1 = 3
	arg2 = 2
	expected = false
	actual, err = le(arg1, arg2)
	if err != nil {
		t.Errorf("le(%v, %v) error: %v", arg1, arg2, err)
		return
	}
	if actual != expected {
		t.Errorf("le(%v, %v) = %v, expected %v", arg1, arg2, actual, expected)
	}
}
