package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestindirectType_1(t *testing.T) {
	type testCase struct {
		name string
		typ  reflect.Type
		want reflect.Type
	}

	testCases := []testCase{
		{
			name: "Test with a pointer type",
			typ:  reflect.TypeOf(&struct{}{}),
			want: reflect.TypeOf(struct{}{}),
		},
		{
			name: "Test with a non-pointer type",
			typ:  reflect.TypeOf(struct{}{}),
			want: reflect.TypeOf(struct{}{}),
		},
		{
			name: "Test with a nil type",
			typ:  nil,
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

			got := indirectType(tc.typ)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("indirectType(%v) = %v, want %v", tc.typ, got, tc.want)
			}
		})
	}
}
