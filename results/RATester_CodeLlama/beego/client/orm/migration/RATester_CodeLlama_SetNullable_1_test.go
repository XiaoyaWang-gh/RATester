package migration

import (
	"fmt"
	"testing"
)

func TestSetNullable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Column{}
	c.SetNullable(true)
	if c.Null != "" {
		t.Errorf("SetNullable failed")
	}
	c.SetNullable(false)
	if c.Null != "NOT NULL" {
		t.Errorf("SetNullable failed")
	}
}
