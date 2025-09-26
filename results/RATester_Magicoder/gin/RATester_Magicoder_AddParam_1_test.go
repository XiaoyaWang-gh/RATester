package gin

import (
	"fmt"
	"testing"
)

func TestAddParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Params: make(Params, 0),
	}

	key := "testKey"
	value := "testValue"

	ctx.AddParam(key, value)

	if len(ctx.Params) != 1 {
		t.Errorf("Expected length of Params to be 1, but got %d", len(ctx.Params))
	}

	param := ctx.Params[0]
	if param.Key != key || param.Value != value {
		t.Errorf("Expected param to be %s:%s, but got %s:%s", key, value, param.Key, param.Value)
	}
}
