package utils

import (
	"fmt"
	"testing"
)

func TestGetDisplayString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := []interface{}{"hello", "world"}
	expected := "hello world"
	actual := GetDisplayString(data...)
	if actual != expected {
		t.Errorf("GetDisplayString() = %v, want %v", actual, expected)
	}
}
