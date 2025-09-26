package orm

import (
	"fmt"
	"testing"
)

func TestValues_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &rawSet{
		query: "SELECT * FROM table",
		args:  []interface{}{},
		orm:   &ormBase{},
	}

	container := &[]Params{}
	cols := []string{"col1", "col2"}

	count, err := o.Values(container, cols...)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if count == 0 {
		t.Errorf("Expected count to be greater than 0, got %v", count)
	}

	if len(*container) != int(count) {
		t.Errorf("Expected length of container to be %v, got %v", count, len(*container))
	}

	for _, row := range *container {
		for _, col := range cols {
			if _, ok := row[col]; !ok {
				t.Errorf("Expected row to contain column %v, got %v", col, row)
			}
		}
	}
}
