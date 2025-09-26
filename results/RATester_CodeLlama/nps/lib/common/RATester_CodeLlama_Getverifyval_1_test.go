package common

import (
	"fmt"
	"testing"
)

func TestGetverifyval_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vkey := "123456"
	want := "123456"
	got := Getverifyval(vkey)
	if got != want {
		t.Errorf("Getverifyval() = %v, want %v", got, want)
	}
}
