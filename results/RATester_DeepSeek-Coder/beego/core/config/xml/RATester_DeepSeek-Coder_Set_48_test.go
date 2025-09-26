package xml

import (
	"fmt"
	"testing"
)

func TestSet_48(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: make(map[string]interface{}),
	}

	key := "testKey"
	val := "testVal"

	err := cc.Set(key, val)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if cc.data[key] != val {
		t.Errorf("Expected value to be %v, got %v", val, cc.data[key])
	}
}
