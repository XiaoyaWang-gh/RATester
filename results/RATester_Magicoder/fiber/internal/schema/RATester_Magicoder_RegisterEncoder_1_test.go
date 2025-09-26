package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRegisterEncoder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Encoder{
		regenc: make(map[reflect.Type]encoderFunc),
	}

	type testStruct struct {
		Field string
	}

	encoder := func(v reflect.Value) string {
		return v.Field(0).String()
	}

	e.RegisterEncoder(testStruct{}, encoder)

	ts := testStruct{"test"}
	expected := "test"
	actual := e.regenc[reflect.TypeOf(ts)](reflect.ValueOf(ts))

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
