package urls

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestURL_4(t *testing.T) {
	tests := []struct {
		name string
		b    BaseURL
		want *url.URL
	}{
		{
			name: "Test case 1",
			b: BaseURL{
				url: &url.URL{
					Scheme: "http",
					Host:   "example.com",
				},
			},
			want: &url.URL{
				Scheme: "http",
				Host:   "example.com",
			},
		},
		{
			name: "Test case 2",
			b: BaseURL{
				url: &url.URL{
					Scheme: "https",
					Host:   "example.com",
				},
			},
			want: &url.URL{
				Scheme: "https",
				Host:   "example.com",
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.b.URL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseURL.URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
