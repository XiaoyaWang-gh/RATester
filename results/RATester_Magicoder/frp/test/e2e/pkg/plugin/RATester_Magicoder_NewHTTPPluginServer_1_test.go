package plugin

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/test/e2e/mock/server/httpserver"
)

func TestNewHTTPPluginServer_1(t *testing.T) {
	type args struct {
		port      int
		newFunc   NewPluginRequest
		handler   Handler
		tlsConfig *tls.Config
	}
	tests := []struct {
		name string
		args args
		want *httpserver.Server
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

			if got := NewHTTPPluginServer(tt.args.port, tt.args.newFunc, tt.args.handler, tt.args.tlsConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPPluginServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
