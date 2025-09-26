package httplib

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDoRequestWithCtx_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		b        *BeegoHTTPRequest
		args     args
		wantResp *http.Response
		wantErr  bool
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

			gotResp, err := tt.b.DoRequestWithCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeegoHTTPRequest.DoRequestWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("BeegoHTTPRequest.DoRequestWithCtx() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
