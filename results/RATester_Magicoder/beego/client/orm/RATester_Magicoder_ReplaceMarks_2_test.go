package orm

import (
	"fmt"
	"testing"
)

func TestReplaceMarks_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBasePostgres{}
	query := "SELECT * FROM table WHERE col1 = ? AND col2 = ?"
	db.ReplaceMarks(&query)
	expected := "SELECT * FROM table WHERE col1 = $1 AND col2 = $2"
	if query != expected {
		t.Errorf("Expected %s, but got %s", expected, query)
	}
}
