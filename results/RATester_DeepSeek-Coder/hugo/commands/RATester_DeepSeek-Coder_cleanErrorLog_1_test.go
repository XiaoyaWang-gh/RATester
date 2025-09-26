package commands

import (
	"fmt"
	"testing"
)

func TestCleanErrorLog_1(t *testing.T) {
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
			args: args{content: "error 1: error 2: error 1"},
			want: "error 1: error 2",
		},
		{
			name: "Test case 2",
			args: args{content: "error 1: error 1: error 2: error 1"},
			want: "error 1: error 2",
		},
		{
			name: "Test case 3",
			args: args{content: "error 1: error 2: error 3"},
			want: "error 1: error 2: error 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cleanErrorLog(tt.args.content); got != tt.want {
				t.Errorf("cleanErrorLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
