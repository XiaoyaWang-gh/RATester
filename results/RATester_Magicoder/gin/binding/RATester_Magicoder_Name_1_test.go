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

	toml := tomlBinding{}
	name := toml.Name()
	if name != "toml" {
		t.Errorf("Expected 'toml', but got '%s'", name)
	}
}
