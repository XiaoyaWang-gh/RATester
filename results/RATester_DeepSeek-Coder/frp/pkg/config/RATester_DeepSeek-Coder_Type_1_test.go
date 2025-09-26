package config

import (
	"fmt"
	"testing"
)

func TestType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	flag := &BoolFuncFlag{}
	expected := "bool"
	actual := flag.Type()
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
