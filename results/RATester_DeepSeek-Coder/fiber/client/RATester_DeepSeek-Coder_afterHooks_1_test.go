package client

import (
	"fmt"
	"testing"
)

func TestAfterHooks_1(t *testing.T) {
	type args struct {
		resp *Response
	}
	tests := []struct {
		name    string
		c       *core
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

			if err := tt.c.afterHooks(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("core.afterHooks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
