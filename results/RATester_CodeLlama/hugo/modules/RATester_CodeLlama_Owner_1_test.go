package modules

import (
	"fmt"
	"testing"
)

func TestOwner_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{}
	if m.Owner() != nil {
		t.Errorf("Owner() = %v, want nil", m.Owner())
	}
}
