package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func Testinit_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	funcs := make(map[string]reflect.Value)
	unFuncs := make(map[string]bool)

	funcs["init"] = reflect.ValueOf(v)
	unFuncs["init"] = true

	reflect.ValueOf(v).Method(0).Call(nil)

	for i := 0; i < reflect.TypeOf(v).NumMethod(); i++ {
		m := reflect.TypeOf(v).Method(i)
		if !unFuncs[m.Name] {
			if _, ok := funcs[m.Name]; !ok {
				t.Errorf("Method %s not found in funcs map", m.Name)
			}
		}
	}
}
