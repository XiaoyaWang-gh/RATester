package testhelpers

import (
	"fmt"
	"net/url"
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
			name: "Test with valid URL",
			args: args{rawURL: "https://www.example.com"},
			want: &url.URL{
				Scheme: "https",
				Host:   "www.example.com",
			},
		},
		{
			name: "Test with invalid URL",
			args: args{rawURL: ":com"},
			want: &url.URL{
				Scheme: "",
				Host:   "",
				Path:   ":com",
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

			got := MustParseURL(tt.args.rawURL)
			if got.Scheme != tt.want.Scheme || got.Host != tt.want.Host || got.Path != tt.want.Path {
				t.Errorf("MustParseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
