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
			name: "Test with string value",
			a:    Attribute{Name: "test", Value: "test_value"},
			want: "test_value",
		},
		{
			name: "Test with int value",
			a:    Attribute{Name: "test", Value: 123},
			want: "123",
		},
		{
			name: "Test with float value",
			a:    Attribute{Name: "test", Value: 123.456},
			want: "123.456",
		},
		{
			name: "Test with bool value",
			a:    Attribute{Name: "test", Value: true},
			want: "true",
		},
		{
			name: "Test with nil value",
			a:    Attribute{Name: "test", Value: nil},
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
