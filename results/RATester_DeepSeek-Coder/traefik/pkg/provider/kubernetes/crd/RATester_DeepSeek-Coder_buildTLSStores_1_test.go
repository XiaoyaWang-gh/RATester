package crd

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tls"
)

func TestBuildTLSStores_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		client Client
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]tls.Store
		want1 map[string]*tls.CertAndStores
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

			got, got1 := buildTLSStores(tt.args.ctx, tt.args.client)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTLSStores() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("buildTLSStores() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
