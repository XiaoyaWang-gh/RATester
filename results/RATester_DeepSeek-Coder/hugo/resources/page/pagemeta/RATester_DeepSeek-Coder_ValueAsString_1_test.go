package pagemeta

import (
	"fmt"
	"testing"
)

func TestValueAsString_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    Source
		want string
	}{
		{
			name: "nil value",
			s:    Source{Value: nil},
			want: "",
		},
		{
			name: "string value",
			s:    Source{Value: "test"},
			want: "test",
		},
		{
			name: "int value",
			s:    Source{Value: 123},
			want: "123",
		},
		{
			name: "float value",
			s:    Source{Value: 123.456},
			want: "123.456",
		},
		{
			name: "bool value",
			s:    Source{Value: true},
			want: "true",
		},
	}

	for _, tt := range tests {
		tt := tt // NOTE: https://golang.org/doc/faq#closures_and_goroutines
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			if got := tt.s.ValueAsString(); got != tt.want {
				t.Errorf("Source.ValueAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
