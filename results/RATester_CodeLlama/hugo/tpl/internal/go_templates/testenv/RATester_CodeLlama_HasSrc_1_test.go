package testenv

import (
	"fmt"
	"testing"
)

func TestHasSrc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if HasSrc() {
		t.Error("HasSrc() should return false")
	}
}
