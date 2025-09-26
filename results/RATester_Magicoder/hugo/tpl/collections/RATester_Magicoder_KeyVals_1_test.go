package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKeyVals_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	key := "testKey"
	values := []any{1, 2, 3}

	result, err := ns.KeyVals(key, values...)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if result.Key != key {
		t.Errorf("Expected key to be %v, but got: %v", key, result.Key)
	}

	if !reflect.DeepEqual(result.Values, values) {
		t.Errorf("Expected values to be %v, but got: %v", values, result.Values)
	}
}
