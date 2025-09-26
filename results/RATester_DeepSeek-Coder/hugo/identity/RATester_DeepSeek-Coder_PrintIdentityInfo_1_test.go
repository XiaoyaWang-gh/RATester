package identity

import (
	"fmt"
	"testing"
)

func TestPrintIdentityInfo_1(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test with a string",
			args: args{v: "test"},
		},
		{
			name: "Test with a slice of int",
			args: args{v: []int{1, 2, 3}},
		},
		{
			name: "Test with a struct",
			args: args{v: struct{ Name string }{Name: "test"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			PrintIdentityInfo(tt.args.v)
		})
	}
}
