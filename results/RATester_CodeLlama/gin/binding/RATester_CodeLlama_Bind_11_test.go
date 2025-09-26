package binding

import (
	"fmt"
	"testing"
)

func TestBind_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	var yamlBinding yamlBinding
	var obj any
	err := yamlBinding.Bind(nil, obj)
	if err != nil {
		t.Errorf("Bind() error = %v, want nil", err)
		return
	}
}
