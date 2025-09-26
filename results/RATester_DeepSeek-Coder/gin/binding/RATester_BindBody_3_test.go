package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindBody_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := msgpackBinding{}

	type testStruct struct {
		Field1 string `msgpack:"field1"`
		Field2 int    `msgpack:"field2"`
	}

	testBody := []byte{0x83, 0xa5, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0xa5, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0xc3, 0x85, 0x49, 0x01}
	testObj := &testStruct{}

	err := b.BindBody(testBody, testObj)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedObj := &testStruct{
		Field1: "field1",
		Field2: 1,
	}

	if !reflect.DeepEqual(testObj, expectedObj) {
		t.Errorf("Expected %v, got %v", expectedObj, testObj)
	}
}
