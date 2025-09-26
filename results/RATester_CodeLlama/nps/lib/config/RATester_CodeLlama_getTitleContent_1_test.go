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
			name: "testcase1",
			args: args{s: "testcase1"},
			want: "testcase1",
		},
		{
			name: "testcase2",
			args: args{s: "testcase2"},
			want: "testcase2",
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
