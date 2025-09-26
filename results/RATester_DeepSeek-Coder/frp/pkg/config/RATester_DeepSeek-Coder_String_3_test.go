package config

import (
	"fmt"
	"testing"
)

func TestString_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	flag := &BoolFuncFlag{
		v: true,
	}

	expected := "true"
	actual := flag.String()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}

	flag.v = false
	expected = "false"
	actual = flag.String()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
