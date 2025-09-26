package server

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"
)

func TestGetGetConfigForClient_1(t *testing.T) {
	tests := []struct {
		name    string
		e       *http3server
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

			got, err := tt.e.getGetConfigForClient(tt.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("getGetConfigForClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGetConfigForClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
