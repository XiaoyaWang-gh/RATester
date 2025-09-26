package server

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"
)

func TestGetGetConfigForClient_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		server  *http3server
		info    *tls.ClientHelloInfo
		want    *tls.Config
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

			got, err := tt.server.getGetConfigForClient(tt.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("http3server.getGetConfigForClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("http3server.getGetConfigForClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
