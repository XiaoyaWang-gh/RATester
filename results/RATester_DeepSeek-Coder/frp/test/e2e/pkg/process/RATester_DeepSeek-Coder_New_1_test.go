package process

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestNew_1(t *testing.T) {
	type args struct {
		path   string
		params []string
	}
	tests := []struct {
		name string
		args args
		want *Process
	}{
		{
			name: "TestNew",
			args: args{
				path:   "test",
				params: []string{"param1", "param2"},
			},
			want: &Process{
				cmd: &exec.Cmd{
					Path: "test",
					Args: []string{"test", "param1", "param2"},
				},
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

			got := New(tt.args.path, tt.args.params)
			if !reflect.DeepEqual(got.cmd, tt.want.cmd) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
