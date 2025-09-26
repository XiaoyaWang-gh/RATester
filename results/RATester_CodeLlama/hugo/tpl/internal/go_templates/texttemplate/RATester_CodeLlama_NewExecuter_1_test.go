package template

import (
	"fmt"
	"testing"
)

func TestNewExecuter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var helper ExecHelper
	executer := NewExecuter(helper)
	if executer == nil {
		t.Errorf("NewExecuter() = nil, want not nil")
	}
}
