package util

import (
	"fmt"
	"testing"
)

func TestConstantTimeEqString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if !ConstantTimeEqString("a", "a") {
		t.Errorf("ConstantTimeEqString(\"a\", \"a\") = false, want true")
	}
	if ConstantTimeEqString("a", "b") {
		t.Errorf("ConstantTimeEqString(\"a\", \"b\") = true, want false")
	}
	if ConstantTimeEqString("", "") {
		t.Errorf("ConstantTimeEqString(\"\", \"\") = true, want false")
	}
}
