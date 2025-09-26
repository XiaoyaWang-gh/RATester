package context

import (
	"fmt"
	"testing"
)

func TestJSONResp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Output: &BeegoOutput{},
	}

	data := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}

	err := ctx.JSONResp(data)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
