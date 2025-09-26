package codegen

import (
	"fmt"
	"testing"
)

func TestfirstToUpper_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{name: "hello"},
			want: "Hello",
		},
		{
			name: "Test Case 2",
			args: args{name: "world"},
			want: "World",
		},
		{
			name: "Test Case 3",
			args: args{name: "golang"},
			want: "Golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := firstToUpper(tt.args.name); got != tt.want {
				t.Errorf("firstToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
