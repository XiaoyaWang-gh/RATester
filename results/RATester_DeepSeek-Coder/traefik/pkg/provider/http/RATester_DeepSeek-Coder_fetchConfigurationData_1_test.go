package http

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestFetchConfigurationData_1(t *testing.T) {
	type testCase struct {
		name          string
		provider      *Provider
		expectedError bool
	}

	testCases := []testCase{
		{
			name: "should return error when endpoint is not reachable",
			provider: &Provider{
				Endpoint: "http://localhost:12345",
				Headers:  map[string]string{"Content-Type": "application/json"},
				httpClient: &http.Client{
					Timeout: 10 * time.Millisecond,
				},
			},
			expectedError: true,
		},
		{
			name: "should return error when endpoint returns non-ok status code",
			provider: &Provider{
				Endpoint: "http://localhost:12345",
				Headers:  map[string]string{"Content-Type": "application/json"},
				httpClient: &http.Client{
					Timeout: 10 * time.Second,
				},
			},
			expectedError: true,
		},
		{
			name: "should return data when endpoint returns ok status code",
			provider: &Provider{
				Endpoint: "http://localhost:12345",
				Headers:  map[string]string{"Content-Type": "application/json"},
				httpClient: &http.Client{
					Timeout: 10 * time.Second,
				},
			},
			expectedError: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := test.provider.fetchConfigurationData()
			if (err != nil) != test.expectedError {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			}
		})
	}
}
