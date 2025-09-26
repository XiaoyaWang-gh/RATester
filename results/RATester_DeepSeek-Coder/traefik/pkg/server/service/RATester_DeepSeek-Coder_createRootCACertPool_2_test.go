package service

import (
	"crypto/x509"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestCreateRootCACertPool_2(t *testing.T) {
	type args struct {
		rootCAs []types.FileOrContent
	}
	tests := []struct {
		name string
		args args
		want *x509.CertPool
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

			if got := createRootCACertPool(tt.args.rootCAs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createRootCACertPool() = %v, want %v", got, tt.want)
			}
		})
	}
}
