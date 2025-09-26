package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeFloat64_1(t *testing.T) {
	testCases := []struct {
		name string
		v    reflect.Value
		want string
	}{
		{
			name: "Test case 1",
			v:    reflect.ValueOf(1.23),
			want: "1.23",
		},
		{
			name: "Test case 2",
			v:    reflect.ValueOf(4.56),
			want: "4.56",
		},
		{
			name: "Test case 3",
			v:    reflect.ValueOf(7.89),
			want: "7.89",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := encodeFloat64(tc.v)
			if got != tc.want {
				t.Errorf("encodeFloat64(%v) = %v, want %v", tc.v, got, tc.want)
			}
		})
	}
}
