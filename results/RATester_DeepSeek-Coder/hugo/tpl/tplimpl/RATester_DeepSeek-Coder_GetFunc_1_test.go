package tplimpl

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGetFunc_1(t *testing.T) {
	ctx := context.Background()
	helper := &templateExecHelper{
		funcs: map[string]reflect.Value{
			"testFunc": reflect.ValueOf(func(ctx context.Context) {}),
		},
	}

	t.Run("should return function and context as first argument when function has context as first argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fn, firstArg, found := helper.GetFunc(ctx, nil, "testFunc")
		if !found {
			t.Fatal("function not found")
		}
		if fn.Kind() != reflect.Func {
			t.Errorf("expected kind to be Func, got %v", fn.Kind())
		}
		if firstArg.Interface() != ctx {
			t.Errorf("expected first argument to be context, got %v", firstArg.Interface())
		}
	})

	t.Run("should return function and zero value as first argument when function has no context as first argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fn, firstArg, found := helper.GetFunc(ctx, nil, "nonExistentFunc")
		if found {
			t.Errorf("expected function not to be found, but it was")
		}
		if fn.Kind() != reflect.Invalid {
			t.Errorf("expected kind to be Invalid, got %v", fn.Kind())
		}
		if !firstArg.IsValid() {
			t.Errorf("expected first argument to be valid")
		}
	})
}
