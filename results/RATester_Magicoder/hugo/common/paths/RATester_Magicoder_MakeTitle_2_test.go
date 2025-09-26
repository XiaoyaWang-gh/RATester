package paths

import (
	"fmt"
	"testing"
)

func TestMakeTitle_2(t *testing.T) {
	type args struct {
		inpath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				inpath: "hello-world",
			},
			want: "hello world",
		},
		{
			name: "Test Case 2",
			args: args{
				inpath: "test-case-2",
			},
			want: "test case 2",
		},
		{
			name: "Test Case 3",
			args: args{
				inpath: "test-case-3-with-multiple-dashes",
			},
			want: "test case 3 with multiple dashes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MakeTitle(tt.args.inpath); got != tt.want {
				t.Errorf("MakeTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
