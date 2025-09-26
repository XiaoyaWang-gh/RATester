package page

import (
	"fmt"
	"testing"
)

func TestTrimSpace_1(t *testing.T) {
	type args struct {
		ss string
	}
	tests := []struct {
		name string
		s    HtmlSummary
		args args
		want string
	}{
		{
			name: "Test with spaces",
			s:    HtmlSummary{},
			args: args{
				ss: "   Hello, World   ",
			},
			want: "Hello, World",
		},
		{
			name: "Test with no spaces",
			s:    HtmlSummary{},
			args: args{
				ss: "Hello, World",
			},
			want: "Hello, World",
		},
		{
			name: "Test with special characters",
			s:    HtmlSummary{},
			args: args{
				ss: "   Hello, @World   #$%^   ",
			},
			want: "Hello, @World   #$%^",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.trimSpace(tt.args.ss); got != tt.want {
				t.Errorf("HtmlSummary.trimSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
