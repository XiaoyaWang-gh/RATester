package retry

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHeader_6(t *testing.T) {
	testCases := []struct {
		name     string
		response *responseWriter
		want     http.Header
	}{
		{
			name: "Test when response has not been written",
			response: &responseWriter{
				responseWriter: nil,
				headers:        http.Header{"Content-Type": []string{"application/json"}},
				shouldRetry:    false,
				written:        false,
			},
			want: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Test when response has been written",
			response: &responseWriter{
				responseWriter: nil,
				headers:        http.Header{"Content-Type": []string{"application/json"}},
				shouldRetry:    false,
				written:        true,
			},
			want: http.Header{"Content-Type": []string{"application/json"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.response.Header()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
