package main

import (
	"fmt"
	"testing"
)

func TestWriteln_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ew := &errWriter{}
	a := []interface{}{"a", "b", "c"}
	ew.writeln(a...)
	if ew.err != nil {
		t.Errorf("writeln() error = %v, wantErr %v", ew.err, nil)
	}
}
