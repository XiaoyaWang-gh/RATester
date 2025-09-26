package acme

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestSetConfigListenerChan_1(t *testing.T) {
	testCases := []struct {
		desc             string
		chanSize         int
		expectedChanSize int
	}{
		{
			desc:             "Test SetConfigListenerChan with a channel of size 10",
			chanSize:         10,
			expectedChanSize: 10,
		},
		{
			desc:             "Test SetConfigListenerChan with a channel of size 0",
			chanSize:         0,
			expectedChanSize: 0,
		},
		{
			desc:             "Test SetConfigListenerChan with a channel of size 100",
			chanSize:         100,
			expectedChanSize: 100,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			provider := &Provider{}
			configChan := make(chan dynamic.Configuration, test.chanSize)
			provider.SetConfigListenerChan(configChan)

			if len(configChan) != test.expectedChanSize {
				t.Errorf("Expected channel size to be %d, got %d", test.expectedChanSize, len(configChan))
			}
		})
	}
}
