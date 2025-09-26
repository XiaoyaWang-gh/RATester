package orm

import (
	"fmt"
	"testing"
)

func TestRaw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	query := "SELECT * FROM table"
	args := []interface{}{}
	rs := o.Raw(query, args...)
	if rs == nil {
		t.Errorf("Expected RawSeter, got nil")
	}
}
