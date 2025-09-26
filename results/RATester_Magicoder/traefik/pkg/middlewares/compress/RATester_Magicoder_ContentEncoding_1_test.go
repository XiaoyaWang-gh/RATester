package compress

import (
	"fmt"
	"testing"
)

func TestContentEncoding_1(t *testing.T) {
	tests := []struct {
		name string
		alg  string
		want string
	}{
		{
			name: "Test case 1",
			alg:  "gzip",
			want: "gzip",
		},
		{
			name: "Test case 2",
			alg:  "deflate",
			want: "deflate",
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

			c := &compressionWriter{
				alg: tt.alg,
			}
			if got := c.ContentEncoding(); got != tt.want {
				t.Errorf("ContentEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
