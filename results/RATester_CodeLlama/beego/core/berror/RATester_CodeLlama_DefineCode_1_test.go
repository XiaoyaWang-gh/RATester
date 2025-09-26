package berror

import (
	"fmt"
	"testing"
)

func TestDefineCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	code := uint32(1)
	module := "module"
	name := "name"
	desc := "desc"
	res := DefineCode(code, module, name, desc)
	if res.Code() != code {
		t.Errorf("code is not equal, want %d, got %d", code, res.Code())
	}
	if res.Module() != module {
		t.Errorf("module is not equal, want %s, got %s", module, res.Module())
	}
	if res.Desc() != desc {
		t.Errorf("desc is not equal, want %s, got %s", desc, res.Desc())
	}
	if res.Name() != name {
		t.Errorf("name is not equal, want %s, got %s", name, res.Name())
	}
}
