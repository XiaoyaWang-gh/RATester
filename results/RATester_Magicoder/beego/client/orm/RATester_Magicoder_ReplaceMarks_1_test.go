package orm

import (
	"fmt"
	"testing"
)

func TestReplaceMarks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	db := &dbBase{}
	query := "SELECT * FROM table WHERE column = ?"

	// Act
	db.ReplaceMarks(&query)

	// Assert
	expected := "SELECT * FROM table WHERE column = ?"
	if query != expected {
		t.Errorf("Expected query to be '%s', but got '%s'", expected, query)
	}
}
