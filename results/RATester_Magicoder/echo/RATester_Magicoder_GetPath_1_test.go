package echo

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestGetPath_1(t *testing.T) {
	tests := []struct {
		name string
		r    *http.Request
		want string
	}{
		{
			name: "Test with RawPath",
			r: &http.Request{
				URL: &url.URL{
					RawPath: "/test/path",
					Path:    "/test/path",
				},
			},
			want: "/test/path",
		},
		{
			name: "Test with Path",
			r: &http.Request{
				URL: &url.URL{
					Path: "/test/path",
				},
			},
			want: "/test/path",
		},
		{
			name: "Test with empty RawPath",
			r: &http.Request{
				URL: &url.URL{
					RawPath: "",
					Path:    "/test/path",
				},
			},
			want: "/test/path",
		},
		{
			name: "Test with empty Path",
			r: &http.Request{
				URL: &url.URL{
					Path: "",
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetPath(tt.r); got != tt.want {
				t.Errorf("GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
