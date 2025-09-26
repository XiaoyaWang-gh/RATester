package validation

import (
	"fmt"
	"testing"
)

func TestMergeParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	obj := "test"
	params := []interface{}{"test"}
	result := mergeParam(v, obj, params)
	if len(result) != 3 {
		t.Errorf("mergeParam failed")
	}
}
