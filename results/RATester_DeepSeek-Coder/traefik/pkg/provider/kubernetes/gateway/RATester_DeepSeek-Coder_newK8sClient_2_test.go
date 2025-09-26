package gateway

import (
	"context"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewK8sClient_2(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		provider      *Provider
		expectedError bool
	}{
		{
			name: "in-cluster client",
			provider: &Provider{
				LabelSelector: "test",
			},
			expectedError: false,
		},
		{
			name: "cluster-external client",
			provider: &Provider{
				Endpoint:      "https://localhost:6443",
				LabelSelector: "test",
			},
			expectedError: false,
		},
		{
			name: "invalid label selector",
			provider: &Provider{
				LabelSelector: "test=",
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := context.Background()
			client, err := tc.provider.newK8sClient(ctx)
			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
			}
		})
	}
}
