package util

import (
	"fmt"
	"testing"
)

func TestConstantTimeEqString_1(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				a: "test1",
				b: "test1",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				a: "test2",
				b: "test3",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				a: "",
				b: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ConstantTimeEqString(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ConstantTimeEqString() = %v, want %v", got, tt.want)
			}
		})
	}
}
