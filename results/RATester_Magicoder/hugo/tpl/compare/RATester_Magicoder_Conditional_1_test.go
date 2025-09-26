package compare

import (
	"fmt"
	"testing"
)

func TestConditional_1(t *testing.T) {
	n := &Namespace{}

	tests := []struct {
		name   string
		cond   bool
		v1     any
		v2     any
		expect any
	}{
		{
			name:   "true",
			cond:   true,
			v1:     "v1",
			v2:     "v2",
			expect: "v1",
		},
		{
			name:   "false",
			cond:   false,
			v1:     "v1",
			v2:     "v2",
			expect: "v2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := n.Conditional(test.cond, test.v1, test.v2)
			if result != test.expect {
				t.Errorf("Expected %v, but got %v", test.expect, result)
			}
		})
	}
}
