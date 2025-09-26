package template

import (
	"fmt"
	"testing"
)

func TestOption_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	template := &Template{
		name: "test",
	}

	options := []string{"opt1", "opt2"}
	result := template.Option(options...)

	if result.name != "test" {
		t.Errorf("Expected name to be 'test', but got %s", result.name)
	}
}
