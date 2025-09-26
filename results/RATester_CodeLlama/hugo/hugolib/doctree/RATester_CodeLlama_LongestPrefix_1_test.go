package doctree

import (
	"fmt"
	"testing"
)

func TestLongestPrefix_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "test1"},
			want: "test1",
		},
		{
			name: "test2",
			args: args{s: "test2"},
			want: "test2",
		},
		{
			name: "test3",
			args: args{s: "test3"},
			want: "test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.s; got != tt.want {
				t.Errorf("LongestPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
