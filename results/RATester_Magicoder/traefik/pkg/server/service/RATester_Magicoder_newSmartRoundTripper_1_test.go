package service

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestnewSmartRoundTripper_1(t *testing.T) {
	transport := &http.Transport{}
	forwardingTimeouts := &dynamic.ForwardingTimeouts{}

	tests := []struct {
		name               string
		transport          *http.Transport
		forwardingTimeouts *dynamic.ForwardingTimeouts
		wantErr            bool
	}{
		{
			name:               "Test case 1",
			transport:          transport,
			forwardingTimeouts: forwardingTimeouts,
			wantErr:            false,
		},
		{
			name:               "Test case 2",
			transport:          nil,
			forwardingTimeouts: forwardingTimeouts,
			wantErr:            true,
		},
		{
			name:               "Test case 3",
			transport:          transport,
			forwardingTimeouts: nil,
			wantErr:            false,
		},
		{
			name:               "Test case 4",
			transport:          nil,
			forwardingTimeouts: nil,
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := newSmartRoundTripper(tt.transport, tt.forwardingTimeouts)
			if (err != nil) != tt.wantErr {
				t.Errorf("newSmartRoundTripper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
