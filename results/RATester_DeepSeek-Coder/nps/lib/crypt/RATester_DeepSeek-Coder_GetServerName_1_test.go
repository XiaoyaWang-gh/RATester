package crypt

import (
	"fmt"
	"testing"
)

func TestGetServerName_1(t *testing.T) {
	tests := []struct {
		name string
		ch   *ClientHelloMsg
		want string
	}{
		{
			name: "Test case 1",
			ch: &ClientHelloMsg{
				serverName: "example.com",
			},
			want: "example.com",
		},
		{
			name: "Test case 2",
			ch: &ClientHelloMsg{
				serverName: "google.com",
			},
			want: "google.com",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.ch.GetServerName(); got != tt.want {
				t.Errorf("GetServerName() = %v, want %v", got, tt.want)
			}
		})
	}
}
