package binding

import (
	"fmt"
	"testing"
)

func TestName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	binding := tomlBinding{}
	if binding.Name() != "toml" {
		t.Errorf("binding.Name() = %v, want %v", binding.Name(), "toml")
	}
}
