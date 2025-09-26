package limit

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"golang.org/x/time/rate"
)

func TestNewReader_1(t *testing.T) {
	type args struct {
		r       io.Reader
		limiter *rate.Limiter
	}
	tests := []struct {
		name string
		args args
		want *Reader
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

			if got := NewReader(tt.args.r, tt.args.limiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
