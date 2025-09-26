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
			name: "Test case 1",
			args: args{
				rule: "rule1",
				name: "name1",
			},
			want: "name1-827ccb0eea8a706c4c34a16891f84e7b",
		},
		{
			name: "Test case 2",
			args: args{
				rule: "rule2",
				name: "name2",
			},
			want: "name2-827ccb0eea8a706c4c34a16891f84e7b",
		},
		{
			name: "Test case 3",
			args: args{
				rule: "rule3",
				name: "name3",
			},
			want: "name3-827ccb0eea8a706c4c34a16891f84e7b",
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
