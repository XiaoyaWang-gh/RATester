package attributes

import (
	"fmt"
	"testing"
)

func TestValueString_1(t *testing.T) {
	tests := []struct {
		name string
		a    Attribute
		want string
	}{
		{
			name: "string",
			a:    Attribute{Name: "name", Value: "value"},
			want: "value",
		},
		{
			name: "int",
			a:    Attribute{Name: "name", Value: 1},
			want: "1",
		},
		{
			name: "float",
			a:    Attribute{Name: "name", Value: 1.1},
			want: "1.1",
		},
		{
			name: "bool",
			a:    Attribute{Name: "name", Value: true},
			want: "true",
		},
		{
			name: "nil",
			a:    Attribute{Name: "name", Value: nil},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.a.ValueString(); got != tt.want {
				t.Errorf("ValueString() = %v, want %v", got, tt.want)
			}
		})
	}
}
