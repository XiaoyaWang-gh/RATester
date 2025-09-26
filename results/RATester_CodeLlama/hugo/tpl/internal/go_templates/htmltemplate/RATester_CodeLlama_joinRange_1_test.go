package template

import (
	"fmt"
	"testing"
)

func TestJoinRange_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c0 := context{}
	rc := &rangeContext{}
	c0 = joinRange(c0, rc)
	if c0.state != stateError {
		t.Errorf("joinRange(%v, %v) = %v; want stateError", c0, rc, c0)
	}
}
