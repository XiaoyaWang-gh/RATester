package berror

import (
	"fmt"
	"testing"
)

func TestDefineCode_1(t *testing.T) {
	t.Run("TestDefineCode", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		code := uint32(100)
		module := "testModule"
		name := "testName"
		desc := "testDesc"

		res := DefineCode(code, module, name, desc)

		if res.Code() != code {
			t.Errorf("Expected code %d, got %d", code, res.Code())
		}

		if res.Module() != module {
			t.Errorf("Expected module %s, got %s", module, res.Module())
		}

		if res.Name() != name {
			t.Errorf("Expected name %s, got %s", name, res.Name())
		}

		if res.Desc() != desc {
			t.Errorf("Expected desc %s, got %s", desc, res.Desc())
		}
	})
}
