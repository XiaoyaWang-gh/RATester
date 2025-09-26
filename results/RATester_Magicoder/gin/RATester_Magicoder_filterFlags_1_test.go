package gin

import (
	"fmt"
	"testing"
)

func TestfilterFlags_1(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				content: "field name string",
			},
			want: "field",
		},
		{
			name: "Test case 2",
			args: args{
				content: "field content string",
			},
			want: "field",
		},
		{
			name: "Test case 3",
			args: args{
				content: "field want string",
			},
			want: "field",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := filterFlags(tt.args.content); got != tt.want {
				t.Errorf("filterFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
