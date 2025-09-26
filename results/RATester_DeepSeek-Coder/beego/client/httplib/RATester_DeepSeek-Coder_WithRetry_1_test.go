package httplib

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestWithRetry_1(t *testing.T) {
	type args struct {
		times int
		delay time.Duration
	}
	tests := []struct {
		name string
		args args
		want BeegoHTTPRequestOption
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

			if got := WithRetry(tt.args.times, tt.args.delay); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRetry() = %v, want %v", got, tt.want)
			}
		})
	}
}
