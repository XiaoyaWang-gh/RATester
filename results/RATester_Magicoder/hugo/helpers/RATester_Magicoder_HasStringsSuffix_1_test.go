package helpers

import (
	"fmt"
	"testing"
)

func TestHasStringsSuffix_1(t *testing.T) {
	type args struct {
		s      []string
		suffix []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				s:      []string{"a", "b", "c"},
				suffix: []string{"b", "c"},
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				s:      []string{"a", "b", "c"},
				suffix: []string{"a", "b", "c", "d"},
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				s:      []string{"a", "b", "c"},
				suffix: []string{},
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				s:      []string{},
				suffix: []string{"a", "b", "c"},
			},
			want: false,
		},
		{
			name: "Test case 5",
			args: args{
				s:      []string{},
				suffix: []string{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasStringsSuffix(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("HasStringsSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
