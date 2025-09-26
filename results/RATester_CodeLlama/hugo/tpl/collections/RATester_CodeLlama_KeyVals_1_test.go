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
	key := "key"
	values := []any{"value1", "value2"}
	kvs, err := ns.KeyVals(key, values...)
	if err != nil {
		t.Fatal(err)
	}
	if kvs.Key != key {
		t.Errorf("expected key %q, got %q", key, kvs.Key)
	}
	if !reflect.DeepEqual(kvs.Values, values) {
		t.Errorf("expected values %q, got %q", values, kvs.Values)
	}
}
