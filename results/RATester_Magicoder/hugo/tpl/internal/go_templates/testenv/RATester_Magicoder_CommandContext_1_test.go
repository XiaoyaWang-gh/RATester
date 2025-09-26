package testenv

import (
	"context"
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestCommandContext_1(t *testing.T) {
	type args struct {
		ctx  context.Context
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

			if got := CommandContext(t, tt.args.ctx, tt.args.name, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
