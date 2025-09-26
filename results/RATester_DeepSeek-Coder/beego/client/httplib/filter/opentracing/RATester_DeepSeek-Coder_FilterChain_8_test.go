package opentracing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/httplib"
)

func TestFilterChain_8(t *testing.T) {
	type args struct {
		next httplib.Filter
		ctx  context.Context
		req  *httplib.BeegoHTTPRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
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

			builder := &FilterChainBuilder{}
			got, err := builder.FilterChain(tt.args.next)(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterChain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
