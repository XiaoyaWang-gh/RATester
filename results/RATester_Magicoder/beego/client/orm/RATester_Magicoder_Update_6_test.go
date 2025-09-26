package orm

import (
	"fmt"
	"testing"
)

func TestUpdate_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &struct{}{}
	cols := []string{"col1", "col2"}

	_, err := o.Update(md, cols...)
	if err != nil {
		t.Errorf("Update() error = %v, wantErr %v", err, nil)
		return
	}
}
