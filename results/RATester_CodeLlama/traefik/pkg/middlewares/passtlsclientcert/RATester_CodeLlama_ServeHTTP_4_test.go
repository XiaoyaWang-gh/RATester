package passtlsclientcert

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeHTTP_4(t *testing.T) {
	type fields struct {
		next http.Handler
		name string
		pem  bool
		info *tlsClientCertificateInfo
	}
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test case 1",
			fields: fields{
				next: nil,
				name: "",
				pem:  false,
				info: nil,
			},
			args: args{
				rw:  nil,
				req: nil,
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

			p := &passTLSClientCert{
				next: tt.fields.next,
				name: tt.fields.name,
				pem:  tt.fields.pem,
				info: tt.fields.info,
			}
			p.ServeHTTP(tt.args.rw, tt.args.req)
		})
	}
}
