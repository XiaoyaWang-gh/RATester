package mock

import (
	"fmt"
	"testing"
)

func TestOne_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	d := &DoNothingQuerySetter{}
	var container interface{}
	cols := []string{"column1", "column2"}

	// Act
	err := d.One(container, cols...)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
