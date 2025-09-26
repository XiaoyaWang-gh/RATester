package stripprefix

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_13(t *testing.T) {
	type fields struct {
		next       http.Handler
		prefixes   []string
		name       string
		forceSlash bool
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

			s := &stripPrefix{
				next:       tt.fields.next,
				prefixes:   tt.fields.prefixes,
				name:       tt.fields.name,
				forceSlash: tt.fields.forceSlash,
			}
			got, got1, got2 := s.GetTracingInformation()
			if got != tt.want {
				t.Errorf("stripPrefix.GetTracingInformation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("stripPrefix.GetTracingInformation() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("stripPrefix.GetTracingInformation() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
