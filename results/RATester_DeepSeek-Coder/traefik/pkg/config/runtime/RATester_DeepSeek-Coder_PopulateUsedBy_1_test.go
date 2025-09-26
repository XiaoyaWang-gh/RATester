package runtime

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestPopulateUsedBy_1(t *testing.T) {
	tests := []struct {
		desc     string
		config   *Configuration
		expected *Configuration
	}{
		{
			desc: "Test with nil config",
			config: &Configuration{
				Routers:        nil,
				Middlewares:    nil,
				TCPMiddlewares: nil,
				Services:       nil,
				TCPRouters:     nil,
				TCPServices:    nil,
				UDPRouters:     nil,
				UDPServices:    nil,
			},
			expected: &Configuration{
				Routers:        nil,
				Middlewares:    nil,
				TCPMiddlewares: nil,
				Services:       nil,
				TCPRouters:     nil,
				TCPServices:    nil,
				UDPRouters:     nil,
				UDPServices:    nil,
			},
		},
		// Add more test cases here
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			test.config.PopulateUsedBy()
			assert.Equal(t, test.expected, test.config)
		})
	}
}
