package hreflect

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestIsContextType_1(t *testing.T) {
	testCases := []struct {
		name string
		tp   reflect.Type
		want bool
	}{
		{
			name: "context.Context",
			tp:   reflect.TypeOf((*context.Context)(nil)).Elem(),
			want: true,
		},
		{
			name: "*context.Context",
			tp:   reflect.TypeOf((*context.Context)(nil)),
			want: true,
		},
		{
			name: "context.Context interface",
			tp:   reflect.TypeOf((*interface{})(nil)).Elem(),
			want: true,
		},
		{
			name: "*context.Context interface",
			tp:   reflect.TypeOf((*interface{})(nil)),
			want: true,
		},
		{
			name: "non-context type",
			tp:   reflect.TypeOf(1),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := IsContextType(tc.tp)
			if got != tc.want {
				t.Errorf("IsContextType(%v) = %v, want %v", tc.tp, got, tc.want)
			}
		})
	}
}
