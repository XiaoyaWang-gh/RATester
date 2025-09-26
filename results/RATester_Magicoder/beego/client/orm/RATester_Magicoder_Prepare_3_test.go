package orm

import (
	"fmt"
	"testing"
)

func TestPrepare_3(t *testing.T) {
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

	preparer, err := o.Prepare()
	if err != nil {
		t.Errorf("Error preparing raw set: %v", err)
	}

	if preparer == nil {
		t.Error("Preparer is nil")
	}
}
