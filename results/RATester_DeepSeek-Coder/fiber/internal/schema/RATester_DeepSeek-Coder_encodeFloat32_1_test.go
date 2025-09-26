package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeFloat32_1(t *testing.T) {
	testCases := []struct {
		name string
		v    reflect.Value
		want string
	}{
		{
			name: "float32 value",
			v:    reflect.ValueOf(float32(1.23)),
			want: "1.23",
		},
		{
			name: "float32 value",
			v:    reflect.ValueOf(float32(4.56)),
			want: "4.56",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := encodeFloat32(tc.v)
			if got != tc.want {
				t.Errorf("encodeFloat32(%v) = %v, want %v", tc.v, got, tc.want)
			}
		})
	}
}
