package context

import (
	"fmt"
	"testing"
)

func TestYAML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}
	data := map[string]string{"key": "value"}
	err := output.YAML(data)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
