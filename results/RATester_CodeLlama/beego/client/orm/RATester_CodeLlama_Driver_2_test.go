package orm

import (
	"fmt"
	"testing"
)

func TestDriver_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingOrm{}
	if d.Driver() != nil {
		t.Errorf("Driver() should return nil")
	}
}
