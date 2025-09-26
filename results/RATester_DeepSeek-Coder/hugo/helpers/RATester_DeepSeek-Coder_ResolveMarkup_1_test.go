package helpers

import (
	"fmt"
	"testing"
)

func TestResolveMarkup_1(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		c    *ContentSpec
		args args
		want string
	}{
		{
			name: "Test case 1",
			c:    &ContentSpec{},
			args: args{
				in: "test",
			},
			want: "expected result",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.ResolveMarkup(tt.args.in); got != tt.want {
				t.Errorf("ContentSpec.ResolveMarkup() = %v, want %v", got, tt.want)
			}
		})
	}
}
