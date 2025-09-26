package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSetKeyValueBytes_1(t *testing.T) {
	t.Parallel()

	type args struct {
		host  string
		key   []byte
		value []byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				host:  "example.com",
				key:   []byte("key"),
				value: []byte("value"),
			},
		},
		{
			name: "Test case 2",
			args: args{
				host:  "example.org",
				key:   []byte("another_key"),
				value: []byte("another_value"),
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
				hostCookies: make(map[string][]*fasthttp.Cookie),
			}

			cj.SetKeyValueBytes(tt.args.host, tt.args.key, tt.args.value)

			cookies := cj.getCookiesByHost(tt.args.host)
			if len(cookies) != 1 {
				t.Errorf("Expected 1 cookie, got %d", len(cookies))
			}

			cookie := cookies[0]
			if string(cookie.Key()) != string(tt.args.key) {
				t.Errorf("Expected key %s, got %s", tt.args.key, cookie.Key())
			}

			if string(cookie.Value()) != string(tt.args.value) {
				t.Errorf("Expected value %s, got %s", tt.args.value, cookie.Value())
			}
		})
	}
}
