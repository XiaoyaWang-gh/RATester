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
			name: "Test case 1",
			args: args{
				path:   "/path/to/process",
				params: []string{"param1", "param2"},
			},
			want: &Process{
				cmd: &exec.Cmd{
					Path: "/path/to/process",
					Args: []string{"param1", "param2"},
				},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := New(tt.args.path, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
