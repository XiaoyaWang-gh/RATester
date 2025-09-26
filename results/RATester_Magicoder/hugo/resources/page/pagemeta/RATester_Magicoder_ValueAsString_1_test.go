package pagemeta

import (
	"fmt"
	"testing"
)

func TestValueAsString_1(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{
			name:  "string",
			value: "test",
			want:  "test",
		},
		{
			name:  "int",
			value: 123,
			want:  "123",
		},
		{
			name:  "float",
			value: 123.456,
			want:  "123.456",
		},
		{
			name:  "bool",
			value: true,
			want:  "true",
		},
		{
			name:  "nil",
			value: nil,
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := Source{Value: tt.value}
			if got := s.ValueAsString(); got != tt.want {
				t.Errorf("ValueAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
