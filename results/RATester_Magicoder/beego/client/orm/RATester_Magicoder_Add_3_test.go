package orm

import (
	"fmt"
	"testing"
)

func TestAdd_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &queryM2M{}
	mds := []interface{}{"test"}
	_, err := o.Add(mds...)
	if err != nil {
		t.Errorf("Add() error = %v, wantErr %v", err, nil)
		return
	}
}
