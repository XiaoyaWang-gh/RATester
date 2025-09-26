package client

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestTLSConfig_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		want *tls.Config
	}{
		{
			name: "Test with default TLS Config",
			want: &tls.Config{
				MinVersion: tls.VersionTLS12,
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

			c := &Client{
				fasthttp: &fasthttp.Client{},
			}

			got := c.TLSConfig()

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TLSConfig() = %v, want %v", got, tc.want)
			}
		})
	}
}
