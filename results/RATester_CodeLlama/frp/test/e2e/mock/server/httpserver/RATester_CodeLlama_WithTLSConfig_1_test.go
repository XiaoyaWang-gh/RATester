package httpserver

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestWithTLSConfig_1(t *testing.T) {
	type args struct {
		tlsConfig *tls.Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				tlsConfig: &tls.Config{},
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

			WithTLSConfig(tt.args.tlsConfig)
		})
	}
}
