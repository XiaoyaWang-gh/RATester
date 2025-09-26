package opentelemetry

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/httplib"
	"go.opentelemetry.io/otel/trace"
)

func TestNewOpenTelemetryFilter_1(t *testing.T) {
	type args struct {
		tagURL   bool
		spanFunc CustomSpanFunc
	}
	tests := []struct {
		name string
		args args
		want *OtelFilterChainBuilder
	}{
		{
			name: "test case 1",
			args: args{
				tagURL: true,
				spanFunc: func(span trace.Span, ctx context.Context, req *httplib.BeegoHTTPRequest, resp *http.Response, err error) {
					// TODO: test span
				},
			},
			want: &OtelFilterChainBuilder{
				tagURL: true,
				customSpanFunc: func(span trace.Span, ctx context.Context, req *httplib.BeegoHTTPRequest, resp *http.Response, err error) {
					// TODO: test span
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewOpenTelemetryFilter(tt.args.tagURL, tt.args.spanFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOpenTelemetryFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
