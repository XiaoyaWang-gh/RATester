package gateway

import (
	"context"
	"fmt"
	"testing"
)

func TestNewK8sClient_2(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name          string
		provider      *Provider
		expectedError error
	}{
		{
			name: "valid label selector",
			provider: &Provider{
				LabelSelector: "app=myapp",
			},
			expectedError: nil,
		},
		{
			name: "invalid label selector",
			provider: &Provider{
				LabelSelector: "app=myapp,",
			},
			expectedError: fmt.Errorf("invalid label selector: %q", "app=myapp,"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := tc.provider.newK8sClient(ctx)
			if err != nil && tc.expectedError == nil {
				t.Errorf("expected no error, but got: %v", err)
			}
			if err == nil && tc.expectedError != nil {
				t.Errorf("expected error: %v, but got no error", tc.expectedError)
			}
			if err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("expected error: %v, but got: %v", tc.expectedError, err)
			}
		})
	}
}
