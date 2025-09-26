package testenv

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestCommand_3(t *testing.T) {
	type args struct {
		name string
		args []string
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Command(t, tt.args.name, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Command() = %v, want %v", got, tt.want)
			}
		})
	}
}
