package client

import (
	"fmt"
	"testing"
)

func TestSetFileName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var n string
	var f *File
	SetFileName(n)(f)
	if f.name != n {
		t.Errorf("SetFileName failed")
	}
}
