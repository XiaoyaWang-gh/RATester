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

	// Create a new Migration and Column
	m := &Migration{}
	c := &Column{Name: "test"}

	// Call the SetPrimary method
	c = c.SetPrimary(m)

	// Check that the Column was added to the Primary field of the Migration
	if len(m.Primary) != 1 || m.Primary[0] != c {
		t.Errorf("SetPrimary failed. Expected %v, got %v", c, m.Primary[0])
	}
}
