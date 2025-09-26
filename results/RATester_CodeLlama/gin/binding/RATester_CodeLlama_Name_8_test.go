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

	yamlBinding := yamlBinding{}
	if yamlBinding.Name() != "yaml" {
		t.Errorf("yamlBinding.Name() = %v, want %v", yamlBinding.Name(), "yaml")
	}
}
