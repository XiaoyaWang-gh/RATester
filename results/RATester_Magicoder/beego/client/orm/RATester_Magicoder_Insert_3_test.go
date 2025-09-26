package orm

import (
	"fmt"
	"testing"
)

func TestInsert_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &struct{}{}
	_, err := o.Insert(md)
	if err != nil {
		t.Errorf("Insert() error = %v, wantErr %v", err, nil)
		return
	}
}
