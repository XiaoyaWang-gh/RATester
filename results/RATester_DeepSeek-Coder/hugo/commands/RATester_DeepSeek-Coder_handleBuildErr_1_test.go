package commands

import (
	"fmt"
	"testing"
)

func TestHandleBuildErr_1(t *testing.T) {
	type fields struct {
		r        *rootCommand
		errState hugoBuilderErrState
	}
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			c := &hugoBuilder{
				r:        tt.fields.r,
				errState: tt.fields.errState,
			}
			c.handleBuildErr(tt.args.err, tt.args.msg)
		})
	}
}
