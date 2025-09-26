package migration

import (
	"fmt"
	"testing"
)

func TestSetUnsigned_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Column{}
	c.SetUnsigned(true)
	if c.Unsign != "UNSIGNED" {
		t.Errorf("SetUnsigned failed")
	}
}
