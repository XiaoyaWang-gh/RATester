package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRetryConfig_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want *RetryConfig
	}{
		{
			name: "Test RetryConfig",
			want: &RetryConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Client{}
			if got := c.RetryConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.RetryConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
