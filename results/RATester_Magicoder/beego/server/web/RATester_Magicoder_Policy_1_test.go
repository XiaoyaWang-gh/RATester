package web

import (
	"fmt"
	"testing"
)

func TestPolicy_1(t *testing.T) {
	type args struct {
		pattern string
		method  string
		policy  []PolicyFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestPolicy_0",
			args: args{
				pattern: "/test",
				method:  "GET",
				policy:  []PolicyFunc{},
			},
		},
		{
			name: "TestPolicy_1",
			args: args{
				pattern: "/test",
				method:  "POST",
				policy:  []PolicyFunc{},
			},
		},
		{
			name: "TestPolicy_2",
			args: args{
				pattern: "/test",
				method:  "PUT",
				policy:  []PolicyFunc{},
			},
		},
		{
			name: "TestPolicy_3",
			args: args{
				pattern: "/test",
				method:  "DELETE",
				policy:  []PolicyFunc{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Policy(tt.args.pattern, tt.args.method, tt.args.policy...)
		})
	}
}
