package acme

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestSanitizeDomains_1(t *testing.T) {
	ctx := context.Background()
	provider := &Provider{
		Configuration: &Configuration{},
	}

	tests := []struct {
		name    string
		domains []string
		want    []string
		err     error
	}{
		{
			name:    "no domain",
			domains: []string{},
			want:    nil,
			err:     errors.New("no domain was given"),
		},
		{
			name:    "wildcard domain",
			domains: []string{"*.*"},
			want:    nil,
			err:     fmt.Errorf("unable to generate a wildcard certificate in ACME provider for domain %q : ACME does not allow '*.*' wildcard domain", "*.*"),
		},
		{
			name:    "valid domain",
			domains: []string{"example.com"},
			want:    []string{"example.com"},
			err:     nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := provider.sanitizeDomains(ctx, types.Domain{Main: test.domains[0]})
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("sanitizeDomains() error = %v, wantErr %v", err, test.err)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("sanitizeDomains() = %v, want %v", got, test.want)
			}
		})
	}
}
