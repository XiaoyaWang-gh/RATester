package prometheus

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/httplib"
)

func TestReport_2(t *testing.T) {
	type args struct {
		startTime time.Time
		endTime   time.Time
		ctx       context.Context
		req       *httplib.BeegoHTTPRequest
		resp      *http.Response
		err       error
	}
	tests := []struct {
		name string
		args args
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
			builder.report(tt.args.startTime, tt.args.endTime, tt.args.ctx, tt.args.req, tt.args.resp, tt.args.err)
		})
	}
}
