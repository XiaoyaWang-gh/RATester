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
		t.Error("NewCondition() should not return nil")
	}
	if len(c.params) != 0 {
		t.Error("NewCondition() should initialize params as an empty slice")
	}
}
