package client

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/utils/v2"
)

func TestSetJSONUnmarshal_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		f    utils.JSONUnmarshal
		want utils.JSONUnmarshal
	}{
		{
			name: "TestSetJSONUnmarshal",
			f: func(data []byte, v any) error {
				return json.Unmarshal(data, v)
			},
			want: func(data []byte, v any) error {
				return json.Unmarshal(data, v)
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

			c := &Client{}
			got := c.SetJSONUnmarshal(tt.f)
			if !reflect.DeepEqual(got.jsonUnmarshal, tt.want) {
				t.Errorf("Client.SetJSONUnmarshal() = %v, want %v", got.jsonUnmarshal, tt.want)
			}
		})
	}
}
