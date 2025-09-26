package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tls"
)

func TestGetTLS_2(t *testing.T) {
	type args struct {
		k8sClient  Client
		secretName string
		namespace  string
	}
	tests := []struct {
		name    string
		args    args
		want    *tls.CertAndStores
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

			got, err := getTLS(tt.args.k8sClient, tt.args.secretName, tt.args.namespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTLS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTLS() = %v, want %v", got, tt.want)
			}
		})
	}
}
