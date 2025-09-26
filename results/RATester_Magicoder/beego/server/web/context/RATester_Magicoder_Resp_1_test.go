package context

import (
	"fmt"
	"testing"
)

func TestResp_1(t *testing.T) {
	type args struct {
		ctx  *Context
		data interface{}
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

			if err := tt.args.ctx.Resp(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Resp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
