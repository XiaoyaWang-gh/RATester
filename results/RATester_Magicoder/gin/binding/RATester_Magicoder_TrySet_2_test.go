package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrySet_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hs := headerSource{"Content-Type": []string{"application/json"}}
	value := reflect.ValueOf(hs)
	field := reflect.StructField{Name: "ContentType", Type: reflect.TypeOf("")}
	tagValue := "Content-Type"
	opt := setOptions{isDefaultExists: false, defaultValue: ""}

	result, err := hs.TrySet(value, field, tagValue, opt)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !result {
		t.Errorf("Expected TrySet to return true, but got false")
	}
}
