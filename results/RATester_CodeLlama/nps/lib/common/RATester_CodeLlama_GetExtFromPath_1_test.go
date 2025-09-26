package common

import (
	"fmt"
	"testing"
)

func TestGetExtFromPath_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				path: "test.txt",
			},
			want: "txt",
		},
		{
			name: "test case 2",
			args: args{
				path: "test.txt.txt",
			},
			want: "txt",
		},
		{
			name: "test case 3",
			args: args{
				path: "test.txt.txt.txt",
			},
			want: "txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetExtFromPath(tt.args.path); got != tt.want {
				t.Errorf("GetExtFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
