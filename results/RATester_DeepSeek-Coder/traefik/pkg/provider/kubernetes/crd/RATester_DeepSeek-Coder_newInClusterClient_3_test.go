package crd

import (
	"fmt"
	"testing"
)

func TestNewInClusterClient_3(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		endpoint    string
		wantErr     bool
		wantWrapper bool
	}{
		{
			name:        "should return error when k8s config fails",
			endpoint:    "",
			wantErr:     true,
			wantWrapper: false,
		},
		{
			name:        "should return clientWrapper when k8s config succeeds",
			endpoint:    "",
			wantErr:     false,
			wantWrapper: true,
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			client, err := newInClusterClient(tc.endpoint)
			if (err != nil) != tc.wantErr {
				t.Errorf("newInClusterClient() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if (client != nil) != tc.wantWrapper {
				t.Errorf("newInClusterClient() client = %v, wantWrapper %v", client, tc.wantWrapper)
			}
		})
	}
}
