package client

import (
	"fmt"
	"testing"
)

func TestPathParam_1(t *testing.T) {
	type test struct {
		name     string
		r        *Request
		key      string
		expected string
	}

	tests := []test{
		{
			name: "TestPathParam_ExistingKey",
			r: &Request{
				path: &PathParam{
					"key1": "value1",
					"key2": "value2",
				},
			},
			key:      "key1",
			expected: "value1",
		},
		{
			name: "TestPathParam_NonExistingKey",
			r: &Request{
				path: &PathParam{
					"key1": "value1",
					"key2": "value2",
				},
			},
			key:      "key3",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.r.PathParam(tt.key)
			if got != tt.expected {
				t.Errorf("PathParam() = %v, want %v", got, tt.expected)
			}
		})
	}
}
