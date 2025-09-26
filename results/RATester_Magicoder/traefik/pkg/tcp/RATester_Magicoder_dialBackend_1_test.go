package tcp

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestDialBackend_1(t *testing.T) {
	type fields struct {
		address       string
		proxyProtocol *dynamic.ProxyProtocol
		dialer        Dialer
	}
	tests := []struct {
		name    string
		fields  fields
		want    WriteCloser
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

			p := Proxy{
				address:       tt.fields.address,
				proxyProtocol: tt.fields.proxyProtocol,
				dialer:        tt.fields.dialer,
			}
			got, err := p.dialBackend()
			if (err != nil) != tt.wantErr {
				t.Errorf("dialBackend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dialBackend() = %v, want %v", got, tt.want)
			}
		})
	}
}
