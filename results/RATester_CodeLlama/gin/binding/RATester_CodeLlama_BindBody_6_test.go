package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var tomlBinding tomlBinding
	var obj = &struct {
		Name string
	}{}
	err := tomlBinding.BindBody([]byte(`
Name = "test"
`), obj)
	if err != nil {
		t.Errorf("BindBody() error = %v", err)
		return
	}
	if obj.Name != "test" {
		t.Errorf("BindBody() obj.Name = %v, want %v", obj.Name, "test")
	}
}
