package binding

import (
	"fmt"
	"testing"
)

func TestName_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	y := yamlBinding{}
	name := y.Name()
	if name != "yaml" {
		t.Errorf("Expected 'yaml', but got '%s'", name)
	}
}
