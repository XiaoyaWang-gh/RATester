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

	testVal := "test"
	testEncoder := func(v reflect.Value) string {
		return v.String() + " encoded"
	}

	e.RegisterEncoder(testVal, testEncoder)

	val := reflect.ValueOf(testVal)
	encoded := e.regenc[val.Type()](val)

	expected := testEncoder(val)
	if encoded != expected {
		t.Errorf("Expected %q, got %q", expected, encoded)
	}
}
