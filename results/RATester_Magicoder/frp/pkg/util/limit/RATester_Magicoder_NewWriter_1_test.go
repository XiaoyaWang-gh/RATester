package limit

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"

	"golang.org/x/time/rate"
)

func TestNewWriter_1(t *testing.T) {
	tests := []struct {
		name    string
		w       io.Writer
		limiter *rate.Limiter
		want    *Writer
	}{
		{
			name:    "Test case 1",
			w:       &bytes.Buffer{},
			limiter: rate.NewLimiter(1, 1),
			want: &Writer{
				w:       &bytes.Buffer{},
				limiter: rate.NewLimiter(1, 1),
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

			if got := NewWriter(tt.w, tt.limiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
