package data

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsHTTPError_1(t *testing.T) {
	tests := []struct {
		name string
		res  *http.Response
		want bool
	}{
		{
			name: "status code less than 200",
			res: &http.Response{
				StatusCode: 199,
			},
			want: true,
		},
		{
			name: "status code greater than 299",
			res: &http.Response{
				StatusCode: 300,
			},
			want: true,
		},
		{
			name: "status code between 200 and 299",
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
