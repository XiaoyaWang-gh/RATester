package context

import (
	"fmt"
	"testing"
)

func TestYamlResp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Input: &BeegoInput{
			Context: &Context{
				Output: &BeegoOutput{},
			},
		},
	}

	data := struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}

	err := ctx.YamlResp(data)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
