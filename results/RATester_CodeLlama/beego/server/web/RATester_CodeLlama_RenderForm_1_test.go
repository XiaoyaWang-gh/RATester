package web

import (
	"fmt"
	"html/template"
	"testing"
)

func TestRenderForm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Name string
		Age  int
	}
	testStruct := TestStruct{
		Name: "test",
		Age:  18,
	}
	result := RenderForm(testStruct)
	if result != template.HTML("") {
		t.Errorf("RenderForm() = %v, want %v", result, template.HTML(""))
	}
}
