package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/utils/v2"
)

func TestSetXMLUnmarshal_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		f    utils.XMLUnmarshal
		want utils.XMLUnmarshal
	}{
		{
			name: "TestSetXMLUnmarshal",
			f: func(data []byte, v any) error {
				return nil
			},
			want: func(data []byte, v any) error {
				return nil
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
			got := c.SetXMLUnmarshal(tt.f)
			if !reflect.DeepEqual(got.xmlUnmarshal, tt.want) {
				t.Errorf("Client.SetXMLUnmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
