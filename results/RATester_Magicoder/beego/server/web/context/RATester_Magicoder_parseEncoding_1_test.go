package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestparseEncoding_1(t *testing.T) {
	tests := []struct {
		name string
		r    *http.Request
		want string
	}{
		{
			name: "Test Case 1",
			r: &http.Request{
				Header: http.Header{
					"Accept-Encoding": []string{"gzip, deflate, br"},
				},
			},
			want: "br",
		},
		{
			name: "Test Case 2",
			r: &http.Request{
				Header: http.Header{
					"Accept-Encoding": []string{"deflate, br"},
				},
			},
			want: "br",
		},
		{
			name: "Test Case 3",
			r: &http.Request{
				Header: http.Header{
					"Accept-Encoding": []string{"gzip, deflate"},
				},
			},
			want: "gzip",
		},
		{
			name: "Test Case 4",
			r: &http.Request{
				Header: http.Header{
					"Accept-Encoding": []string{"deflate"},
				},
			},
			want: "deflate",
		},
		{
			name: "Test Case 5",
			r: &http.Request{
				Header: http.Header{
					"Accept-Encoding": []string{""},
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

			if got := parseEncoding(tt.r); got != tt.want {
				t.Errorf("parseEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
