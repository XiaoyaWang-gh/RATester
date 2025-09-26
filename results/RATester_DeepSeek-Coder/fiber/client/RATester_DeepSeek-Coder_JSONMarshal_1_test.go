package client

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/utils/v2"
)

func TestJSONMarshal_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		client   *Client
		expected utils.JSONMarshal
	}{
		{
			name: "Test JSONMarshal",
			client: &Client{
				jsonMarshal: func(v any) ([]byte, error) {
					return json.Marshal(v)
				},
			},
			expected: func(v any) ([]byte, error) {
				return json.Marshal(v)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.client.JSONMarshal()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
