package data

import (
	"fmt"
	"net/http"
	"testing"
)

func TestisHTTPError_1(t *testing.T) {
	tests := []struct {
		name string
		res  *http.Response
		want bool
	}{
		{
			name: "Test 1: HTTP Error",
			res: &http.Response{
				StatusCode: 404,
			},
			want: true,
		},
		{
			name: "Test 2: Non-HTTP Error",
			res: &http.Response{
				StatusCode: 200,
			},
			want: false,
		},
		{
			name: "Test 3: HTTP Error",
			res: &http.Response{
				StatusCode: 500,
			},
			want: true,
		},
		{
			name: "Test 4: Non-HTTP Error",
			res: &http.Response{
				StatusCode: 200,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isHTTPError(tt.res); got != tt.want {
				t.Errorf("isHTTPError() = %v, want %v", got, tt.want)
			}
		})
	}
}
