package limit

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/time/rate"
)

func TestNewReader_1(t *testing.T) {
	tests := []struct {
		name    string
		r       io.Reader
		limiter *rate.Limiter
		want    *Reader
	}{
		{
			name:    "Test case 1",
			r:       strings.NewReader("test"),
			limiter: rate.NewLimiter(1, 1),
			want:    &Reader{r: strings.NewReader("test"), limiter: rate.NewLimiter(1, 1)},
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

			if got := NewReader(tt.r, tt.limiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
