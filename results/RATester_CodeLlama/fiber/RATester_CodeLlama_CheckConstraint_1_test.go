package fiber

import (
	"fmt"
	"testing"
)

func TestCheckConstraint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Constraint{}
	param := "test"
	if c.CheckConstraint(param) {
		t.Errorf("CheckConstraint() = %v, want %v", true, false)
	}
}
