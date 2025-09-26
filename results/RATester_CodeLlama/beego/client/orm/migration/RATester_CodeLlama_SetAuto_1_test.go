package migration

import (
	"fmt"
	"testing"
)

func TestSetAuto_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Column{}
	c.SetAuto(true)
	if c.Inc != "auto_increment" {
		t.Errorf("SetAuto failed")
	}
}
