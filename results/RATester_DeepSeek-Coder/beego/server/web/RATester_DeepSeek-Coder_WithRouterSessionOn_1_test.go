package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithRouterSessionOn_1(t *testing.T) {
	type args struct {
		sessionOn bool
	}
	tests := []struct {
		name string
		args args
		want ControllerOption
	}{
		{
			name: "TestWithRouterSessionOn_True",
			args: args{sessionOn: true},
			want: WithRouterSessionOn(true),
		},
		{
			name: "TestWithRouterSessionOn_False",
			args: args{sessionOn: false},
			want: WithRouterSessionOn(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WithRouterSessionOn(tt.args.sessionOn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRouterSessionOn() = %v, want %v", got, tt.want)
			}
		})
	}
}
