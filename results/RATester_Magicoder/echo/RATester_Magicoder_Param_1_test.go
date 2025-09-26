package echo

import (
	"fmt"
	"testing"
)

func TestParam_1(t *testing.T) {
	type test struct {
		name     string
		context  *context
		param    string
		expected string
	}

	tests := []test{
		{
			name: "TestParam_0",
			context: &context{
				pnames:  []string{"id"},
				pvalues: []string{"123"},
			},
			param:    "id",
			expected: "123",
		},
		{
			name: "TestParam_1",
			context: &context{
				pnames:  []string{"id"},
				pvalues: []string{"123"},
			},
			param:    "name",
			expected: "",
		},
		{
			name: "TestParam_2",
			context: &context{
				pnames:  []string{"id", "name"},
				pvalues: []string{"123", "John"},
			},
			param:    "name",
			expected: "John",
		},
		{
			name: "TestParam_3",
			context: &context{
				pnames:  []string{"id", "name"},
				pvalues: []string{"123", "John"},
			},
			param:    "age",
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

			result := tt.context.Param(tt.param)
			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}
		})
	}
}
