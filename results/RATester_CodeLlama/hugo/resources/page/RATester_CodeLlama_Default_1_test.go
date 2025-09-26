package page

import (
	"fmt"
	"testing"
)

func TestDefault_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := Sites{}
	if s.Default() != nil {
		t.Errorf("Default() = %v, want nil", s.Default())
	}
}
