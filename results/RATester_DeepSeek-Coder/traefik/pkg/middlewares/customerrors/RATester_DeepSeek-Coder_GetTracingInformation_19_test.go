package customerrors

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_19(t *testing.T) {
	type fields struct {
		name           string
		next           http.Handler
		backendHandler http.Handler
		httpCodeRanges types.HTTPCodeRanges
		backendQuery   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  string
		want2  trace.SpanKind
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

			c := &customErrors{
				name:           tt.fields.name,
				next:           tt.fields.next,
				backendHandler: tt.fields.backendHandler,
				httpCodeRanges: tt.fields.httpCodeRanges,
				backendQuery:   tt.fields.backendQuery,
			}
			got, got1, got2 := c.GetTracingInformation()
			if got != tt.want {
				t.Errorf("customErrors.GetTracingInformation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("customErrors.GetTracingInformation() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("customErrors.GetTracingInformation() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
