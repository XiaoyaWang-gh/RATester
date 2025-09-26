package migration

import (
	"fmt"
	"testing"
)

func TestAddPrimary_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	primary := &Column{}
	m.AddPrimary(primary)
	if len(m.Primary) != 1 {
		t.Errorf("AddPrimary failed")
	}
}
