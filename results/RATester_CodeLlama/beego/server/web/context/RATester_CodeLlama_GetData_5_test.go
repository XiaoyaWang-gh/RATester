package context

import (
	"fmt"
	"testing"
)

func TestGetData_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	input.data = make(map[interface{}]interface{})
	input.data["key"] = "value"
	if input.GetData("key") != "value" {
		t.Errorf("GetData failed")
	}
}
