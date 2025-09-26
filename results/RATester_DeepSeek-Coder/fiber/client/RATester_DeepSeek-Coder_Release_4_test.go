package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestRelease_4(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want *CookieJar
	}{
		{
			name: "Test case 1",
			want: &CookieJar{
				hostCookies: map[string][]*fasthttp.Cookie{
					"example.com": {
						fasthttp.AcquireCookie(),
						fasthttp.AcquireCookie(),
					},
				},
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

			cj := &CookieJar{
				hostCookies: tt.want.hostCookies,
			}

			cj.Release()

			if cj.hostCookies != nil {
				t.Errorf("Expected hostCookies to be nil, got %v", cj.hostCookies)
			}
		})
	}
}
