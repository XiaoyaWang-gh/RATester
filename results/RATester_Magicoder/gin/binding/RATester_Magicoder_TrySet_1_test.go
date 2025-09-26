package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrySet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	form := formSource{"field": []string{"value"}}
	field := reflect.StructField{Name: "Field"}
	tagValue := "form"
	opt := setOptions{isDefaultExists: true, defaultValue: "default"}

	isSet, err := form.TrySet(reflect.ValueOf(form), field, tagValue, opt)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !isSet {
		t.Errorf("Expected TrySet to return true, but got false")
	}
}
