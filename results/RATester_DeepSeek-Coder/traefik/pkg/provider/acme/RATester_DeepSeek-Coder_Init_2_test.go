package acme

import (
	"errors"
	"fmt"
	"testing"
)

func TestInit_2(t *testing.T) {
	type testCase struct {
		name        string
		provider    *Provider
		expectedErr error
	}

	testCases := []testCase{
		{
			name: "success",
			provider: &Provider{
				Configuration: &Configuration{
					Storage: "test",
				},
			},
			expectedErr: nil,
		},
		{
			name: "no storage",
			provider: &Provider{
				Configuration: &Configuration{
					Storage: "",
				},
			},
			expectedErr: errors.New("unable to initialize ACME provider with no storage location for the certificates"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := test.provider.Init()
			if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", test.expectedErr, err)
			}
			if err != test.expectedErr {
				t.Errorf("expected error %v, got %v", test.expectedErr, err)
			}
		})
	}
}
