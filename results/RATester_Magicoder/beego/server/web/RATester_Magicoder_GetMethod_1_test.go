package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerInfo{
		methods: map[string]string{
			"GET":  "Get",
			"POST": "Post",
		},
	}

	expected := map[string]string{
		"GET":  "Get",
		"POST": "Post",
	}

	result := ctrl.GetMethod()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
