package client

import (
	"fmt"
	"testing"
)

func TestReplace_1(t *testing.T) {
	tests := []struct {
		name      string
		oldClient *Client
		newClient *Client
	}{
		{
			name: "Test Replace",
			oldClient: &Client{
				baseURL: "http://old.com",
			},
			newClient: &Client{
				baseURL: "http://new.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			replace := Replace(tt.newClient)

			if defaultClient != tt.newClient {
				t.Errorf("Expected defaultClient to be %v, got %v", tt.newClient, defaultClient)
			}

			replace()

			if defaultClient != tt.oldClient {
				t.Errorf("Expected defaultClient to be %v, got %v", tt.oldClient, defaultClient)
			}
		})
	}
}
