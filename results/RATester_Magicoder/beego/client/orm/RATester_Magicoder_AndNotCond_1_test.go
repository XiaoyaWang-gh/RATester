package orm

import (
	"fmt"
	"testing"
)

func TestAndNotCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new Condition
	c := &Condition{}

	// Create a new Condition to be used as a sub condition
	subCond := &Condition{}

	// Test the AndNotCond method
	result := c.AndNotCond(subCond)

	// Check if the result is as expected
	if result != c {
		t.Errorf("Expected %v, but got %v", c, result)
	}

	// Test the panic case
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	c.AndNotCond(c)
}
