package orm

import (
	"fmt"
	"testing"
)

func TestValuesFlat_3(t *testing.T) {
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

	container := &ParamsList{}
	cols := []string{"column1", "column2"}

	_, err := o.ValuesFlat(container, cols...)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Add more assertions to check the content of container
}
