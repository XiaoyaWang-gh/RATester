package idempotency

import (
	"fmt"
	"testing"
)

func TestMsgsize_2(t *testing.T) {
	tests := []struct {
		name string
		z    *response
		want int
	}{
		{
			name: "Test case 1",
			z: &response{
				Headers: map[string][]string{
					"Content-Type": {"application/json"},
				},
				Body:       []byte("Hello, World!"),
				StatusCode: 200,
			},
			want: 100,
		},
		{
			name: "Test case 2",
			z: &response{
				Headers: map[string][]string{
					"Content-Type": {"text/html"},
				},
				Body:       []byte("<html><body>Hello, World!</body></html>"),
				StatusCode: 200,
			},
			want: 100,
		},
		{
			name: "Test case 3",
			z: &response{
				Headers: map[string][]string{
					"Content-Type": {"application/xml"},
				},
				Body:       []byte("<xml><body>Hello, World!</body></xml>"),
				StatusCode: 200,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.z.Msgsize(); got != tt.want {
				t.Errorf("Msgsize() = %v, want %v", got, tt.want)
			}
		})
	}
}
