package echo

import (
	"fmt"
	"testing"
)

func TestSetParamValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{}
	c.SetParamNames("name", "age")
	c.SetParamValues("john", "30")
	if c.Param("name") != "john" {
		t.Errorf("Expected name to be john but got %s", c.Param("name"))
	}
	if c.Param("age") != "30" {
		t.Errorf("Expected age to be 30 but got %s", c.Param("age"))
	}
}
