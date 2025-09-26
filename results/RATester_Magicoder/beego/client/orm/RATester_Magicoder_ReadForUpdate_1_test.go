package orm

import (
	"fmt"
	"testing"
)

func TestReadForUpdate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &struct{}{}
	err := o.ReadForUpdate(md)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
