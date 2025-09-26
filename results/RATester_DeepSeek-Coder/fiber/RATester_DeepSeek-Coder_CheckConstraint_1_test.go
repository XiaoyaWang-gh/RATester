package fiber

import (
	"fmt"
	"testing"
)

func TestCheckConstraint_1(t *testing.T) {
	type testCase struct {
		name       string
		constraint Constraint
		param      string
		want       bool
	}

	testCases := []testCase{
		{
			name: "Test Case 1",
			constraint: Constraint{
				ID:   noConstraint,
				Name: "test",
				Data: []string{"test"},
			},
			param: "test",
			want:  true,
		},
		{
			name: "Test Case 2",
			constraint: Constraint{
				ID:   noConstraint,
				Name: "test",
				Data: []string{"test"},
			},
			param: "not_test",
			want:  false,
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.constraint.CheckConstraint(tc.param)
			if got != tc.want {
				t.Errorf("CheckConstraint() = %v, want %v", got, tc.want)
			}
		})
	}
}
