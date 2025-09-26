package validation

import (
	"fmt"
	"testing"
)

func TestmergeParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	obj := "test"
	params := []interface{}{"param1", "param2"}
	result := mergeParam(v, obj, params)

	if len(result) != len(params)+2 {
		t.Errorf("Expected length of result to be %d, but got %d", len(params)+2, len(result))
	}

	if result[0] != v {
		t.Errorf("Expected first element of result to be %v, but got %v", v, result[0])
	}

	if result[1] != obj {
		t.Errorf("Expected second element of result to be %v, but got %v", obj, result[1])
	}

	for i, param := range params {
		if result[i+2] != param {
			t.Errorf("Expected %dth element of result to be %v, but got %v", i+2, param, result[i+2])
		}
	}
}
