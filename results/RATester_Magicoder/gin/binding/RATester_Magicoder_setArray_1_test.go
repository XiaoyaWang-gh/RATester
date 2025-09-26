package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsetArray_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		ArrayField []string
	}

	testStruct := TestStruct{}
	value := reflect.ValueOf(&testStruct).Elem()
	field, _ := value.Type().FieldByName("ArrayField")

	vals := []string{"test1", "test2", "test3"}
	err := setArray(vals, value.FieldByName("ArrayField"), field)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(testStruct.ArrayField) != len(vals) {
		t.Errorf("Expected length of ArrayField to be %d, but got: %d", len(vals), len(testStruct.ArrayField))
	}

	for i, v := range vals {
		if testStruct.ArrayField[i] != v {
			t.Errorf("Expected ArrayField[%d] to be %s, but got: %s", i, v, testStruct.ArrayField[i])
		}
	}
}
