package httpserver

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"
)

func TestWithTLSConfig_1(t *testing.T) {
	type args struct {
		tlsConfig *tls.Config
	}
	tests := []struct {
		name string
		args args
		want *Server
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

			if got := WithTLSConfig(tt.args.tlsConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTLSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
