package template

import (
	"context"
	"fmt"
	"io"
	"testing"
)

func TestExecuteWithContext_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		p    Preparer
		wr   io.Writer
		data any
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
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

			e := &executer{}
			err := e.ExecuteWithContext(tt.args.ctx, tt.args.p, tt.args.wr, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("executer.ExecuteWithContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
