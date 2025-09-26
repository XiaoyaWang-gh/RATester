package ecs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfiguration_1(t *testing.T) {
	testCases := []struct {
		desc     string
		instance ecsInstance
		expected configuration
	}{
		{
			desc: "should return the configuration",
			instance: ecsInstance{
				Labels: map[string]string{
					"traefik.ecs.enable": "true",
				},
			},
			expected: configuration{
				Enable: true,
			},
		},
		{
			desc: "should return the default configuration",
			instance: ecsInstance{
				Labels: map[string]string{
					"traefik.ecs.enable": "false",
				},
			},
			expected: configuration{
				Enable: false,
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			p := &Provider{}
			conf, err := p.getConfiguration(test.instance)
			require.NoError(t, err)
			assert.Equal(t, test.expected, conf)
		})
	}
}
