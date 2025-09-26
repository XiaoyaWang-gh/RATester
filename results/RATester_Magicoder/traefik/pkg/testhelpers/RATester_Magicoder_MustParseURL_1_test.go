package testhelpers

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestMustParseURL_1(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name string
		args args
		want *url.URL
	}{
		{
			name: "Test case 1",
			args: args{rawURL: "http://example.com"},
			want: &url.URL{
				Scheme: "http",
				Host:   "example.com",
			},
		},
		{
			name: "Test case 2",
			args: args{rawURL: "https://user:pass@example.com:8080/path?query#fragment"},
			want: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/path",
				RawQuery: "query",
				Fragment: "fragment",
			},
		},
		{
			name: "Test case 3",
			args: args{rawURL: "invalid"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MustParseURL(tt.args.rawURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustParseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
