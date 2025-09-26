package gin

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	type testCase struct {
		name   string
		key    string
		value  any
		expect any
	}

	testCases := []testCase{
		{
			name:   "set string",
			key:    "name",
			value:  "John",
			expect: "John",
		},
		{
			name:   "set int",
			key:    "age",
			value:  30,
			expect: 30,
		},
		{
			name:   "set bool",
			key:    "isActive",
			value:  true,
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{}
			c.Set(tc.key, tc.value)
			if c.Keys[tc.key] != tc.expect {
				t.Errorf("Expected %v, got %v", tc.expect, c.Keys[tc.key])
			}
		})
	}
}
