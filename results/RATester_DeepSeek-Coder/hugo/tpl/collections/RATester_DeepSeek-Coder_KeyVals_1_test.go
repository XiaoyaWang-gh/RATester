package collections

import (
	"fmt"
	"testing"
	"time"

	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/compare"
)

func TestKeyVals_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{
		loc:      time.UTC,
		sortComp: &compare.Namespace{},
		deps:     &deps.Deps{},
	}

	key := "testKey"
	values := []any{"value1", "value2", "value3"}

	kv, err := ns.KeyVals(key, values...)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if kv.Key != key {
		t.Errorf("Expected key to be %v, got %v", key, kv.Key)
	}

	for i, value := range values {
		if kv.Values[i] != value {
			t.Errorf("Expected value at index %d to be %v, got %v", i, value, kv.Values[i])
		}
	}
}
