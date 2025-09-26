package plugin

import (
	"context"
	"fmt"
	"testing"
)

func TestDo_1(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *Request
		res *Response
	}
	tests := []struct {
		name    string
		p       *httpPlugin
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

			if err := tt.p.do(tt.args.ctx, tt.args.r, tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("httpPlugin.do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
