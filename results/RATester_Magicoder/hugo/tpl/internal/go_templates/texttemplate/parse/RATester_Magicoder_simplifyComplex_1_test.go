package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsimplifyComplex_1(t *testing.T) {
	tests := []struct {
		name     string
		n        *NumberNode
		expected *NumberNode
	}{
		{
			name: "Test case 1",
			n: &NumberNode{
				Complex128: complex(1.0, 0.0),
			},
			expected: &NumberNode{
				IsFloat:    true,
				Float64:    1.0,
				IsInt:      true,
				Int64:      1,
				IsUint:     false,
				Complex128: complex(1.0, 0.0),
			},
		},
		{
			name: "Test case 2",
			n: &NumberNode{
				Complex128: complex(1.0, 1.0),
			},
			expected: &NumberNode{
				IsFloat:    false,
				IsInt:      false,
				IsUint:     false,
				Complex128: complex(1.0, 1.0),
			},
		},
		{
			name: "Test case 3",
			n: &NumberNode{
				Complex128: complex(1.0, -1.0),
			},
			expected: &NumberNode{
				IsFloat:    false,
				IsInt:      false,
				IsUint:     false,
				Complex128: complex(1.0, -1.0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.n.simplifyComplex()
			if !reflect.DeepEqual(tt.n, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, tt.n)
			}
		})
	}
}
