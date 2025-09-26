package gin

import (
	"fmt"
	"testing"
)

func TestNameOfFunction_1(t *testing.T) {
	type args struct {
		f any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_case_1",
			args: args{
				f: func() {},
			},
			want: "github.com/your_name/your_repo.TestnameOfFunction.func1",
		},
		{
			name: "test_case_2",
			args: args{
				f: func() {},
			},
			want: "github.com/your_name/your_repo.TestnameOfFunction.func2",
		},
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
