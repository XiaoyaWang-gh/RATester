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
	driver := d.Driver()
	if driver != nil {
		t.Errorf("Expected nil, got %v", driver)
	}
}
