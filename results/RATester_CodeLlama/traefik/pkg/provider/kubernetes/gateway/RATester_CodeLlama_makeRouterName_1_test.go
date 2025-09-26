package gateway

import (
	"fmt"
	"testing"
)

func TestMakeRouterName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rule := "rule"
	name := "name"
	expected := "name-%.10x"
	actual := makeRouterName(rule, name)
	if actual != expected {
		t.Errorf("makeRouterName(%s, %s) = %s, want %s", rule, name, actual, expected)
	}
}
