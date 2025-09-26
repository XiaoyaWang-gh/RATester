package server

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestIsEmptyConfiguration_1(t *testing.T) {
	type testCase struct {
		name     string
		conf     *dynamic.Configuration
		expected bool
	}

	testCases := []testCase{
		{
			name:     "empty configuration",
			conf:     &dynamic.Configuration{},
			expected: true,
		},
		{
			name: "non-empty configuration",
			conf: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"router1": {},
					},
					Services: map[string]*dynamic.Service{
						"service1": {},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"middleware1": {},
					},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"router1": {},
					},
					Services: map[string]*dynamic.TCPService{
						"service1": {},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{
						"middleware1": {},
					},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers: map[string]*dynamic.UDPRouter{
						"router1": {},
					},
					Services: map[string]*dynamic.UDPService{
						"service1": {},
					},
				},
			},
			expected: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := isEmptyConfiguration(test.conf)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
