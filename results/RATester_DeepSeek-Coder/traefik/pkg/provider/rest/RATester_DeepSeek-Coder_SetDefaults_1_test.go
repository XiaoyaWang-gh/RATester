package rest

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestSetDefaults_1(t *testing.T) {
	type testCase struct {
		name     string
		provider *Provider
		expected *Provider
	}

	testCases := []testCase{
		{
			name: "Test with nil configurationChan",
			provider: &Provider{
				Insecure: true,
			},
			expected: &Provider{
				Insecure: true,
			},
		},
		{
			name: "Test with non-nil configurationChan",
			provider: &Provider{
				Insecure:          true,
				configurationChan: make(chan dynamic.Message),
			},
			expected: &Provider{
				Insecure:          true,
				configurationChan: make(chan dynamic.Message),
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			test.provider.SetDefaults()

			if !reflect.DeepEqual(test.provider, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, test.provider)
			}
		})
	}
}
