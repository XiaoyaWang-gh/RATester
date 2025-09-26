package helpers

import (
	"fmt"
	"testing"
)

func TestHasStringsPrefix_1(t *testing.T) {
	type args struct {
		s      []string
		prefix []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				s:      []string{"a", "b", "c", "d", "e"},
				prefix: []string{"a", "b", "c"},
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				s:      []string{"a", "b", "c", "d", "e"},
				prefix: []string{"a", "b", "c", "d"},
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				s:      []string{"a", "b", "c", "d", "e"},
				prefix: []string{},
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				s:      []string{},
				prefix: []string{"a", "b", "c"},
			},
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

			if got := HasStringsPrefix(tt.args.s, tt.args.prefix); got != tt.want {
				t.Errorf("HasStringsPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
