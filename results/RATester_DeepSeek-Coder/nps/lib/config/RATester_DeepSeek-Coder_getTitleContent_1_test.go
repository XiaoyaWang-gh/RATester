package config

import (
	"fmt"
	"testing"
)

func TestGetTitleContent_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{s: "[[test]]"},
			want: "test",
		},
		{
			name: "Test Case 2",
			args: args{s: "test"},
			want: "test",
		},
		{
			name: "Test Case 3",
			args: args{s: "[[test1][test2]]"},
			want: "test1test2",
		},
		{
			name: "Test Case 4",
			args: args{s: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getTitleContent(tt.args.s); got != tt.want {
				t.Errorf("getTitleContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
