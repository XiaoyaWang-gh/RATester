package context

import (
	"fmt"
	"testing"
)

func TestSetData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	key := "testKey"
	val := "testValue"
	input.SetData(key, val)
	if input.data[key] != val {
		t.Errorf("Expected %v, got %v", val, input.data[key])
	}
}
