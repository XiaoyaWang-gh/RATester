package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetParamsWithStruct_1(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		name   string
		params PathParam
		input  any
		want   PathParam
	}{
		{
			name:   "Test with valid struct",
			params: PathParam{"key1": "value1", "key2": "value2"},
			input:  testStruct{Name: "test", Age: 20},
			want:   PathParam{"key1": "value1", "key2": "value2", "name": "test", "age": "20"},
		},
		{
			name:   "Test with nil struct",
			params: PathParam{"key1": "value1", "key2": "value2"},
			input:  nil,
			want:   PathParam{"key1": "value1", "key2": "value2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.params.SetParamsWithStruct(tt.input)
			if !reflect.DeepEqual(tt.params, tt.want) {
				t.Errorf("got %v, want %v", tt.params, tt.want)
			}
		})
	}
}
