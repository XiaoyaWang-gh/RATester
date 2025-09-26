package utils

import (
	"fmt"
	"testing"
)

func TestExist_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := StrTo("")
	if f.Exist() {
		t.Errorf("Exist() = true, want false")
	}

	f = StrTo("1")
	if !f.Exist() {
		t.Errorf("Exist() = false, want true")
	}
}
