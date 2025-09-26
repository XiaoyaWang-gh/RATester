package hreflect

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestCallMethodByName_1(t *testing.T) {
	type TestStruct struct {
		Name string
	}

	testCases := []struct {
		name     string
		ctx      context.Context
		method   string
		target   TestStruct
		expected []reflect.Value
	}{
		{
			name:   "Test case 1",
			ctx:    context.Background(),
			method: "Name",
			target: TestStruct{Name: "Test"},
			expected: []reflect.Value{
				reflect.ValueOf("Test"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			v := reflect.ValueOf(tc.target)
			result := CallMethodByName(tc.ctx, tc.method, v)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
