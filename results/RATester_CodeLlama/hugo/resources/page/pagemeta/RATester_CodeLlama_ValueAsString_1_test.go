package pagemeta

import (
	"fmt"
	"testing"
)

func TestValueAsString_1(t *testing.T) {

	tests := []struct {
		name string
		s    Source
		want string
	}{
		{
			name: "nil",
			s:    Source{},
			want: "",
		},
		{
			name: "string",
			s:    Source{Value: "foo"},
			want: "foo",
		},
		{
			name: "int",
			s:    Source{Value: 42},
			want: "42",
		},
		{
			name: "float",
			s:    Source{Value: 3.14},
			want: "3.14",
		},
		{
			name: "bool",
			s:    Source{Value: true},
			want: "true",
		},
		{
			name: "struct",
			s:    Source{Value: struct{}{}},
			want: "{}",
		},
		{
			name: "map",
			s:    Source{Value: map[string]string{"foo": "bar"}},
			want: "map[foo:bar]",
		},
		{
			name: "slice",
			s:    Source{Value: []string{"foo", "bar"}},
			want: "[foo bar]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.ValueAsString(); got != tt.want {
				t.Errorf("Source.ValueAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
