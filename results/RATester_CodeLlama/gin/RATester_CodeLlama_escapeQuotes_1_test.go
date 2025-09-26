package gin

import (
	"fmt"
	"testing"
)

func TestEscapeQuotes_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testcase1",
			args: args{
				s: "\"",
			},
			want: "\\\"",
		},
		{
			name: "testcase2",
			args: args{
				s: "\"",
			},
			want: "\\\"",
		},
		{
			name: "testcase3",
			args: args{
				s: "\"",
			},
			want: "\\\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := escapeQuotes(tt.args.s); got != tt.want {
				t.Errorf("escapeQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
