package redactor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDoOnStruct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Field1 string `redact:"true"`
		Field2 string `redact:"false"`
		Field3 string
		Field4 *testStruct
	}

	test := &testStruct{
		Field1: "secret1",
		Field2: "secret2",
		Field3: "secret3",
		Field4: &testStruct{
			Field1: "secret4",
			Field2: "secret5",
			Field3: "secret6",
		},
	}

	err := doOnStruct(reflect.ValueOf(test), "redact", true)
	if err != nil {
		t.Errorf("doOnStruct() error = %v", err)
	}

	if test.Field1 != "" {
		t.Errorf("Field1 should be redacted")
	}

	if test.Field2 != "secret2" {
		t.Errorf("Field2 should not be redacted")
	}

	if test.Field3 != "" {
		t.Errorf("Field3 should be redacted")
	}

	if test.Field4.Field1 != "" {
		t.Errorf("Field4.Field1 should be redacted")
	}

	if test.Field4.Field2 != "secret5" {
		t.Errorf("Field4.Field2 should not be redacted")
	}

	if test.Field4.Field3 != "" {
		t.Errorf("Field4.Field3 should be redacted")
	}
}
