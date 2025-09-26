package httplib

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestWithTimeout_1(t *testing.T) {
	type args struct {
		connectTimeout   time.Duration
		readWriteTimeout time.Duration
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

			if got := WithTimeout(tt.args.connectTimeout, tt.args.readWriteTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
