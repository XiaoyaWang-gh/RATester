package hexec

import (
	"fmt"
	"testing"
)

func TestInPath_1(t *testing.T) {
	type args struct {
		binaryName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				binaryName: "ls",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				binaryName: "non-existing-binary",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				binaryName: "ls/more",
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

			if got := InPath(tt.args.binaryName); got != tt.want {
				t.Errorf("InPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
