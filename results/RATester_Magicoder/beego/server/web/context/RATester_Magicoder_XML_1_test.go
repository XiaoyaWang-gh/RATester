package context

import (
	"fmt"
	"testing"
)

func TestXML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{
		Context: &Context{
			ResponseWriter: &Response{},
		},
	}

	data := struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}{
		Name: "John",
		Age:  30,
	}

	err := output.XML(data, true)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
