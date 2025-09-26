package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/provider/acme"
	"github.com/traefik/traefik/v3/pkg/provider/aggregator"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

func TestInitACMEProvider_1(t *testing.T) {
	type args struct {
		c                                           *static.Configuration
		providerAggregator                          *aggregator.ProviderAggregator
		tlsManager                                  *traefiktls.Manager
		httpChallengeProvider, tlsChallengeProvider challenge.Provider
	}
	tests := []struct {
		name string
		args args
		want []*acme.Provider
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

			if got := initACMEProvider(tt.args.c, tt.args.providerAggregator, tt.args.tlsManager, tt.args.httpChallengeProvider, tt.args.tlsChallengeProvider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initACMEProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
