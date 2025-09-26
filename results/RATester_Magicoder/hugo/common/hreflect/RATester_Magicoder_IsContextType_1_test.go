package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsContextType_1(t *testing.T) {
	tests := []struct {
		name string
		tp   reflect.Type
		want bool
	}{
		{
			name: "contextTypeValue",
			tp:   contextTypeValue,
			want: true,
		},
		{
			name: "contextInterface",
			tp:   contextInterface,
			want: true,
		},
		{
			name: "otherType",
			tp:   reflect.TypeOf(0),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsContextType(tt.tp); got != tt.want {
				t.Errorf("IsContextType() = %v, want %v", got, tt.want)
			}
		})
	}
}
