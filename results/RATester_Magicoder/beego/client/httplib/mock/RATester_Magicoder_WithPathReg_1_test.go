package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithPathReg_1(t *testing.T) {
	type args struct {
		pathReg string
	}
	tests := []struct {
		name string
		args args
		want simpleConditionOption
	}{
		{
			name: "Test case 1",
			args: args{
				pathReg: "test",
			},
			want: WithPathReg("test"),
		},
		{
			name: "Test case 2",
			args: args{
				pathReg: "another test",
			},
			want: WithPathReg("another test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WithPathReg(tt.args.pathReg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPathReg() = %v, want %v", got, tt.want)
			}
		})
	}
}
