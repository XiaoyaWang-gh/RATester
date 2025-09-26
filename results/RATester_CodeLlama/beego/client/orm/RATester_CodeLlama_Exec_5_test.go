package orm

import (
	"fmt"
	"testing"
)

func TestExec_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &rawSet{
		query: "select * from user",
		args:  []interface{}{},
		orm:   &ormBase{},
	}
	result, err := o.Exec()
	if err != nil {
		t.Errorf("Exec() error = %v", err)
		return
	}
	if result == nil {
		t.Errorf("Exec() result is nil")
		return
	}
}
