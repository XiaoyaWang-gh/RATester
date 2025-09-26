package hqt

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArgNames_1(t *testing.T) {
	testCases := []struct {
		name     string
		argNames argNames
		want     []string
	}{
		{
			name:     "Test Case 1",
			argNames: argNames{"arg1", "arg2", "arg3"},
			want:     []string{"arg1", "arg2", "arg3"},
		},
		{
			name:     "Test Case 2",
			argNames: argNames{"arg4", "arg5", "arg6"},
			want:     []string{"arg4", "arg5", "arg6"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.argNames.ArgNames()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}
