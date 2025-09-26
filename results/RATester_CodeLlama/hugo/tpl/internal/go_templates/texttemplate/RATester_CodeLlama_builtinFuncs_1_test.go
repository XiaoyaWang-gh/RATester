package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuiltinFuncs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	got := builtinFuncs()
	want := map[string]reflect.Value{
		"builtinFuncs": reflect.ValueOf(builtinFuncs),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("builtinFuncs() = %v, want %v", got, want)
	}
}
