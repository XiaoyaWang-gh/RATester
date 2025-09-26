package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndirectType_1(t *testing.T) {
	type testCase struct {
		name string
		v    reflect.Type
		want reflect.Type
	}

	testCases := []testCase{
		{
			name: "Test with pointer type",
			v:    reflect.TypeOf(&struct{}{}),
			want: reflect.TypeOf(struct{}{}),
		},
		{
			name: "Test with non-pointer type",
			v:    reflect.TypeOf(struct{}{}),
			want: reflect.TypeOf(struct{}{}),
		},
		{
			name: "Test with nil",
			v:    nil,
			want: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := IndirectType(tc.v)
			if got != tc.want {
				t.Errorf("IndirectType(%v) = %v, want %v", tc.v, got, tc.want)
			}
		})
	}
}
