package passtlsclientcert

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestGetTracingInformation_2(t *testing.T) {
	type fields struct {
		next http.Handler
		name string
		pem  bool
		info *tlsClientCertificateInfo
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

			p := &passTLSClientCert{
				next: tt.fields.next,
				name: tt.fields.name,
				pem:  tt.fields.pem,
				info: tt.fields.info,
			}
			got, got1, got2 := p.GetTracingInformation()
			if got != tt.want {
				t.Errorf("passTLSClientCert.GetTracingInformation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("passTLSClientCert.GetTracingInformation() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("passTLSClientCert.GetTracingInformation() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
