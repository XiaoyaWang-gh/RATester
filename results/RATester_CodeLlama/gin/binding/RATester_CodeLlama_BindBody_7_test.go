package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var yamlBinding yamlBinding
	var obj any
	err := yamlBinding.BindBody([]byte(""), &obj)
	if err != nil {
		t.Errorf("BindBody() error = %v, want nil", err)
		return
	}
	if obj != nil {
		t.Errorf("BindBody() obj = %v, want nil", obj)
	}
}
