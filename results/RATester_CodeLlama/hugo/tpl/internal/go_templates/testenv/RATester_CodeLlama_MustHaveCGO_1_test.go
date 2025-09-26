package testenv

import (
	"fmt"
	"testing"
)

func TestMustHaveCGO_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if !HasCGO() {
		t.Skip("skipping test: no cgo")
	}
}
