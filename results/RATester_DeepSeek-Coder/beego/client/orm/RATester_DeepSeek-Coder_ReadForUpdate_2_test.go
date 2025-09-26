package orm

import (
	"fmt"
	"testing"
)

func TestReadForUpdate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	d := &DoNothingOrm{}
	md := &struct{}{}
	cols := []string{"col1", "col2"}

	// Act
	err := d.ReadForUpdate(md, cols...)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
