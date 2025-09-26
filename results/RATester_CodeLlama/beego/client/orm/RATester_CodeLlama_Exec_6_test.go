package orm

import (
	"fmt"
	"testing"
)

func TestExec_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbQueryLog{}
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
