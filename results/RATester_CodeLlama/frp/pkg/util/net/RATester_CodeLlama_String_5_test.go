package net

import (
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ia := &InternalAddr{}
	if ia.String() != "internal" {
		t.Errorf("ia.String() = %s, want %s", ia.String(), "internal")
	}
}
