package context

import (
	"fmt"
	"testing"
)

func TestResp_1(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		ctx     *Context
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

			if err := tt.ctx.Resp(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Context.Resp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
