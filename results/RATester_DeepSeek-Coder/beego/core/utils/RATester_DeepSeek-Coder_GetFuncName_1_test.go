package utils

import (
	"fmt"
	"testing"
)

func TestGetFuncName_1(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestGetFuncName",
			args: args{
				i: func() {},
			},
			want: "github.com/username/repo.TestGetFuncName.func1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetFuncName(tt.args.i); got != tt.want {
				t.Errorf("GetFuncName() = %v, want %v", got, tt.want)
			}
		})
	}
}
