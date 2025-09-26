package consulcatalog

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/consul/api"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestLoadConfiguration_3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testCases := []struct {
		desc          string
		client        *api.Client
		expectedError error
	}{
		{
			desc:          "error when getConsulServicesData fails",
			client:        &api.Client{},
			expectedError: errors.New("getConsulServicesData error"),
		},
		{
			desc:          "success",
			client:        &api.Client{},
			expectedError: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			provider := &Provider{
				client: test.client,
			}

			mockChan := make(chan dynamic.Message)
			err := provider.loadConfiguration(ctx, &connectCert{}, mockChan)

			if err != test.expectedError {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			}
		})
	}
}
