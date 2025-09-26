package orm

import (
	"fmt"
	"testing"
)

func TestNewCondition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := NewCondition()
	if c == nil {
		t.Errorf("NewCondition() should not return nil")
	}
}
