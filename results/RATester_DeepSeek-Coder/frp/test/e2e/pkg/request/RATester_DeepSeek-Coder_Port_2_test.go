package request

import (
	"crypto/tls"
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestPort_2(t *testing.T) {
	type fields struct {
		protocol  string
		addr      string
		port      int
		body      []byte
		timeout   time.Duration
		resolver  *net.Resolver
		method    string
		host      string
		path      string
		headers   map[string]string
		tlsConfig *tls.Config
		authValue string
		proxyURL  string
	}
	type args struct {
		port int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
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

			r := &Request{
				protocol:  tt.fields.protocol,
				addr:      tt.fields.addr,
				port:      tt.fields.port,
				body:      tt.fields.body,
				timeout:   tt.fields.timeout,
				resolver:  tt.fields.resolver,
				method:    tt.fields.method,
				host:      tt.fields.host,
				path:      tt.fields.path,
				headers:   tt.fields.headers,
				tlsConfig: tt.fields.tlsConfig,
				authValue: tt.fields.authValue,
				proxyURL:  tt.fields.proxyURL,
			}
			if got := r.Port(tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Port() = %v, want %v", got, tt.want)
			}
		})
	}
}
