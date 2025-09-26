package yaml

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
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
		t.Errorf("Set() error = %v, wantErr %v", err, false)
		return
	}

	if cc.data[key] != val {
		t.Errorf("Set() failed, expected %v, got %v", val, cc.data[key])
	}
}
