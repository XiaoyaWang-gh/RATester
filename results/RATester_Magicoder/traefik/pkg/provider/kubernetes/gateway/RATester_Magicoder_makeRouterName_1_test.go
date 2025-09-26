package gateway

import (
	"fmt"
	"testing"
)

func TestMakeRouterName_1(t *testing.T) {
	type args struct {
		rule string
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				rule: "test_rule",
				name: "test_name",
			},
			want: "test_name-9f07e7be556eee24",
		},
		{
			name: "Test Case 2",
			args: args{
				rule: "another_rule",
				name: "another_name",
			},
			want: "another_name-9f07e7be556eee24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := makeRouterName(tt.args.rule, tt.args.name); got != tt.want {
				t.Errorf("makeRouterName() = %v, want %v", got, tt.want)
			}
		})
	}
}
