package orm

import (
	"fmt"
	"testing"
)

func TestLimit_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var limit interface{}
	var args []interface{}
	o.Limit(limit, args...)
	if o.limit != 0 {
		t.Errorf("TestLimit failed")
	}
	if o.offset != 0 {
		t.Errorf("TestLimit failed")
	}
	limit = 10
	args = []interface{}{2}
	o.Limit(limit, args...)
	if o.limit != 10 {
		t.Errorf("TestLimit failed")
	}
	if o.offset != 2 {
		t.Errorf("TestLimit failed")
	}
}
