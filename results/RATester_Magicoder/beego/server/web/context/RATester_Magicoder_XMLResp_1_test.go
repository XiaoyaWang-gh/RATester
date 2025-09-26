package context

import (
	"fmt"
	"testing"
)

func TestXMLResp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Input: &BeegoInput{
			Context: &Context{},
		},
		Output: &BeegoOutput{
			Context: &Context{},
		},
	}

	data := struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}{
		Name: "John",
		Age:  30,
	}

	err := ctx.XMLResp(data)
	if err != nil {
		t.Errorf("XMLResp() error = %v", err)
	}
}
