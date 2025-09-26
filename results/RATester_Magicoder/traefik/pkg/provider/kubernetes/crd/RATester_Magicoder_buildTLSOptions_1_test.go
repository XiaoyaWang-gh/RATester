package crd

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tls"
)

func TestBuildTLSOptions_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		client Client
	}
	tests := []struct {
		name string
		args args
		want map[string]tls.Options
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

			if got := buildTLSOptions(tt.args.ctx, tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTLSOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
