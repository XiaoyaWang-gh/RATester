package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestExec_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DB{
		DB: &sql.DB{},
	}
	query := "query"
	args := []interface{}{}
	result, err := d.Exec(query, args...)
	if err != nil {
		t.Errorf("Exec() error = %v", err)
		return
	}
	if result == nil {
		t.Errorf("Exec() result = nil")
		return
	}
}
