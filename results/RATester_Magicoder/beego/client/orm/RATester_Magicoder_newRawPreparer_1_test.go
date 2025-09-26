package orm

import (
	"fmt"
	"testing"
)

func TestnewRawPreparer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rs := &rawSet{
		query: "SELECT * FROM table",
		args:  []interface{}{},
		orm:   &ormBase{},
	}

	preparer, err := newRawPreparer(rs)
	if err != nil {
		t.Errorf("Error creating RawPreparer: %v", err)
	}

	_, err = preparer.Exec()
	if err != nil {
		t.Errorf("Error executing query: %v", err)
	}

	err = preparer.Close()
	if err != nil {
		t.Errorf("Error closing Preparer: %v", err)
	}
}
