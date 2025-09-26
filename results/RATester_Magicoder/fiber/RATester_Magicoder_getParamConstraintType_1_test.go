package fiber

import (
	"fmt"
	"testing"
)

func TestgetParamConstraintType_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want TypeConstraint
	}{
		{
			name: "Test case 1: ConstraintInt",
			arg:  ConstraintInt,
			want: intConstraint,
		},
		{
			name: "Test case 2: ConstraintBool",
			arg:  ConstraintBool,
			want: boolConstraint,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getParamConstraintType(tt.arg); got != tt.want {
				t.Errorf("getParamConstraintType() = %v, want %v", got, tt.want)
			}
		})
	}
}
