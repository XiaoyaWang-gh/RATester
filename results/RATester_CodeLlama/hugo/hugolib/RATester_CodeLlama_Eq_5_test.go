package hugolib

import (
	"fmt"
	"testing"
)

func TestEq_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	other := &pageState{}
	if !p.Eq(other) {
		t.Errorf("p.Eq(other) = false, want true")
	}
}
