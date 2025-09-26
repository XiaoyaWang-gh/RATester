package gin

import (
	"fmt"
	"testing"
)

func TestnameOfFunction_1(t *testing.T) {
	type args struct {
		f any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				f: func() {},
			},
			want: "github.com/your-module-path/your-function-name",
		},
		{
			name: "Test case 2",
			args: args{
				f: func() {},
			},
			want: "github.com/your-module-path/your-function-name",
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

			if got := nameOfFunction(tt.args.f); got != tt.want {
				t.Errorf("nameOfFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
