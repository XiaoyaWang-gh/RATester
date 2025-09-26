package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetRootCertificateFromString_1(t *testing.T) {
	type args struct {
		pem string
	}
	tests := []struct {
		name string
		c    *Client
		args args
		want *Client
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

			if got := tt.c.SetRootCertificateFromString(tt.args.pem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SetRootCertificateFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
