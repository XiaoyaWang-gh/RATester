package limit

import (
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewWriter(tt.w, tt.limiter)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
