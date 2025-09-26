package migration

import (
	"fmt"
	"testing"
)

func TestSetPrimary_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Column{
		Name: "id",
	}
	m := &Migration{}
	c.SetPrimary(m)
	if len(m.Primary) != 1 {
		t.Errorf("SetPrimary failed")
	}
}
